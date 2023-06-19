package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type IMap map[string]interface{}
type Request int

const (
	GET  Request = 1
	POST Request = 2
)

const (
	TelegramApiUrlFormat string = "https://api.telegram.org/bot%s/%s"
	JsonContentType      string = "application/json"
	TelegramParseMode    string = "MarkdownV2"
)

var (
	EmptyPostError error = errors.New("POST request must contain data.")
	UnknownError   error = errors.New("An unknown error has occurred.")
)

type TelegramClient struct {
	Token                 string                          `json:"token"`
	ParseMode             string                          `json:"parse_mode,omitempty"`
	CommandHandlers       map[string]UpdateHandler        `json:"-"`
	CommandDescriptions   map[string]string               `json:"-"`
	CallbackQueryHandlers map[string]CallbackQueryHandler `json:"-"`
}

type TelegramResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}

/*
NOTE: The parseMode parameter is optional and used for additional markdown.
Documentation: https://core.telegram.org/bots/api#formatting-options
*/
func NewTelegramClient(token string, parseMode string) *TelegramClient {
	tg := &TelegramClient{
		Token:                 token,
		CommandHandlers:       make(map[string]UpdateHandler),
		CommandDescriptions:   make(map[string]string),
		CallbackQueryHandlers: make(map[string]CallbackQueryHandler),
	}

	if len(parseMode) > 0 {
		tg.ParseMode = parseMode
	} else {
		tg.ParseMode = TelegramParseMode
	}

	return tg
}

func TelegramClientError(method string, err error) error {
	return fmt.Errorf("TelegramClient '%s' request failed: %w", method, err)
}

func (tc *TelegramClient) SendRequest(requestMethod Request, method string, data *IMap) (*http.Response, error) {
	if requestMethod == POST && data == nil {
		return nil, EmptyPostError
	}
	url := fmt.Sprintf(TelegramApiUrlFormat, tc.Token, method)
	if requestMethod == GET {
		resp, err := http.Get(url)
		if err != nil {
			return nil, TelegramClientError(method, err)
		}
		return resp, nil
	} else if requestMethod == POST {
		postBytes, err := json.Marshal(*data)
		if err != nil {
			return nil, TelegramClientError(method, err)
		}
		resp, err := http.Post(url, JsonContentType, bytes.NewBuffer(postBytes))
		if err != nil {
			return nil, TelegramClientError(method, err)
		}
		return resp, nil
	}
	return nil, TelegramClientError(method, UnknownError)
}

func (tc *TelegramClient) CleanToken() error {
	if success, err := tc.DeleteMyCommands(nil, ""); success && err != nil {
		return err
	}

	if err := tc.DeleteChatMenuButton(nil); err != nil {
		return err
	}

	if err := tc.DeleteWebhook(true); err != nil {
		return err
	}

	return nil
}
