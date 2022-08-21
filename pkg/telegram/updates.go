package telegram

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ooojustin/telego/pkg/utils"
)

// https://core.telegram.org/bots/api#update
type Update struct {
	ID      int     `json:"update_id"`
	Message Message `mapstructure:"message"`
}

type GetUpdatesResponse struct {
	Ok          bool     `json:"ok"`
	Result      []Update `mapstructure:"result"`
	Description string   `json:"description"`
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
					tc.HandleUpdate(update)
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

func (tc *TelegramClient) HandleUpdate(update Update) {
	data := IMap{
		"update_id": update.ID,
		"chat_id":   update.Message.Chat.ID,
		"username":  update.Message.From.Username,
		"text":      update.Message.Text,
	}
	utils.PrettyPrint(data)
}
