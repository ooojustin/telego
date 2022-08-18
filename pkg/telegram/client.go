package telegram

import (
	"errors"
	"fmt"
	"net/http"
)

type Request int

const (
	GET  Request = 1
	POST Request = 2
)

const (
	TelegramApiUrlFormat string = "https://api.telegram.org/bot%s/%s"
)

var (
	EmptyPostError error = errors.New("POST request must contain data.")
	UnknownError   error = errors.New("An unknown error has occurred.")
)

type TelegramClient struct {
	Token string `json:"token"`
}

func NewTelegramClient(token string) *TelegramClient {
	return &TelegramClient{
		Token: token,
	}
}

func (tc *TelegramClient) SendRequest(requestMethod Request, method string, data *map[string]interface{}) (*http.Response, error) {
	if requestMethod == POST && data == nil {
		return nil, EmptyPostError
	}
	url := fmt.Sprintf(TelegramApiUrlFormat, tc.Token, method)
	if requestMethod == GET {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Errorf("TelegramClient 'GET' failed: %w", err)
		}
		return resp, nil
	}
	return nil, UnknownError
}
