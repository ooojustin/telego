package telegram

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/ooojustin/telego/pkg/utils"
	"golang.org/x/exp/maps"
)

type (
	UpdateHandler        func(Update) error
	CallbackQueryHandler func(Update, []string) error
)

var (
	UnhandledUpdateError  error = errors.New("Unhandled update type.")
	UnhandledCommandError error = errors.New("Unhandled command received.")
	BadUpdateError        error = errors.New("Received update with unexpected conditions.")
)

// https://core.telegram.org/bots/api#update
type Update struct {
	ID            int            `json:"update_id" mapstructure:"update_id"`
	Message       *Message       `json:"message,omitempty" mapstructure:"message,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty" mapstructure:"callback_query,omitempty"`
}

func (u Update) GetType() (string, bool) {
	var data IMap
	mapstructure.Decode(u, &data)

	keys := maps.Keys[IMap](data)
	hasID := utils.Contains[string](keys, "update_id")

	if len(keys) != 2 || !hasID {
		return "", false
	}

	updateType := utils.Remove[string](keys, "update_id")[0]
	return updateType, true
}

type GetUpdatesResponse struct {
	TelegramResponse
	Result []Update `json:"result"`
}

// https://core.telegram.org/bots/api#getupdates
func (tc *TelegramClient) GetUpdates(offset int, allowedUpdates []string) ([]Update, error) {
	var resp *http.Response
	var err error

	req := IMap{}

	if len(allowedUpdates) > 0 {
		req["allowed_updates"] = allowedUpdates
	}

	if offset > 0 {
		req["offset"] = offset
	}

	resp, err = tc.SendRequest(POST, "getUpdates", &req)
	if err != nil {
		return []Update{}, fmt.Errorf("GetUpdates failed: %w", err)
	}

	defer resp.Body.Close()

	var data GetUpdatesResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return []Update{}, fmt.Errorf("GetUpdates failed: %w", err)
	}

	if data.Ok {
		return data.Result, nil
	} else if len(data.Description) > 0 {
		err = errors.New(data.Description)
		return []Update{}, fmt.Errorf("GetUpdates failed: %w", err)
	}

	return []Update{}, fmt.Errorf("GetUpdates failed: %w", UnknownError)
}

func (tc *TelegramClient) StartUpdateHandler(interval time.Duration, allowedUpdates []string) {
	go tc.UpdateHandler(interval, allowedUpdates)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func (tc *TelegramClient) UpdateHandler(interval time.Duration, allowedUpdates []string) {
	var offset int
	quit := make(chan bool)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			updates, err := tc.GetUpdates(offset, allowedUpdates)
			if err == nil {
				lastIdx := len(updates) - 1
				for idx, update := range updates {
					if err := tc.HandleUpdate(update); err != nil {
						fmt.Printf("Error handling update %d: %s", update.ID, err)
					}
					if idx == lastIdx {
						offset = update.ID + 1
					}
				}
			} else {
				fmt.Printf("Failed to get updates in UpdateHandler: %s\n", err)
				quit <- true
			}
		}
	}
}

func (tc *TelegramClient) HandleUpdate(update Update) error {
	updateType, ok := update.GetType()
	if !ok {
		return BadUpdateError
	}

	updateHandlers := IMap{
		"message":        tc.HandleMessage,
		"callback_query": tc.HandleCallbackQuery,
	}

	if vfunc, ok := updateHandlers[updateType]; ok {
		funcName := utils.GetFunctionName(vfunc)
		fmt.Printf("Sending update %d to %s.\n", update.ID, funcName)
		return vfunc.((func(Update) error))(update)
	} else {
		return UnhandledUpdateError
	}
}

func (tc *TelegramClient) HandleMessage(update Update) error {
	if strings.HasPrefix(update.Message.Text, "/") {
		command := update.Message.Text[1:]
		if handler, ok := tc.CommandHandlers[command]; ok {
			return handler(update)
		} else {
			return UnhandledCommandError
		}
	}
	return nil
}

func (tc *TelegramClient) HandleCallbackQuery(update Update) error {
	data := update.CallbackQuery.Data
	for patternStr, handler := range tc.CallbackQueryHandlers {
		if pattern, err := regexp.Compile(patternStr); err == nil {
			if groups := pattern.FindStringSubmatch(data); groups != nil {
				return handler(update, groups)
			}
		}
	}
	return nil
}
