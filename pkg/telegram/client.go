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
)

var (
	EmptyPostError error = errors.New("POST request must contain data.")
	UnknownError   error = errors.New("An unknown error has occurred.")
)

type TelegramClient struct {
	Token    string                   `json:"token"`
	Commands map[string]UpdateHandler `json:"-"`
}

type TelegramResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}

func NewTelegramClient(token string) *TelegramClient {
	return &TelegramClient{
		Token:    token,
		Commands: make(map[string]UpdateHandler),
	}
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
