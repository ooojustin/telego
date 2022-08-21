package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
)

// https://core.telegram.org/bots/api#message
type Message struct {
	ID         int    `json:"message_id"`
	From       User   `json:"from"`
	SenderChat Chat   `json:"sender_chat"`
	Date       int    `json:"date"`
	Chat       Chat   `json:"chat"`
	Text       string `json:"text"`
}

type SendMessageResponse struct {
	TelegramResponse
	Result Message `mapstructure:"result"`
}

func (tc *TelegramClient) SendMessage(chat int, text string) (*Message, error) {
	request := IMap{
		"chat_id": chat,
		"text":    text,
	}

	resp, err := tc.SendRequest(POST, "sendMessage", &request)
	if err != nil {
		return nil, fmt.Errorf("SendMessage failed: %w", err)
	}

	defer resp.Body.Close()

	var data SendMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("SendMessage failed: %w", err)
	}

	if data.Ok {
		return &data.Result, nil
	} else if len(data.Description) > 0 {
		err = errors.New(data.Description)
		return nil, fmt.Errorf("SendMessage failed: %w", err)
	}

	return nil, fmt.Errorf("SendMessage failed: %w", UnknownError)
}
