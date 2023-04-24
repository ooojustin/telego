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
	Result Message `json:"result"`
}

func (tc *TelegramClient) SendMessage(chat int, text string, data *IMap) (*Message, error) {
	request := IMap{
		"chat_id": chat,
		"text":    text,
	}

	if data != nil {
		for k, v := range *data {
			request[k] = v
		}
	}

	resp, err := tc.SendRequest(POST, "sendMessage", &request)
	if err != nil {
		return nil, fmt.Errorf("SendMessage failed: %w", err)
	}

	defer resp.Body.Close()

	var smr SendMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("SendMessage failed: %w", err)
	}

	if smr.Ok {
		return &smr.Result, nil
	} else if len(smr.Description) > 0 {
		err = errors.New(smr.Description)
		return nil, fmt.Errorf("SendMessage failed: %w", err)
	}

	return nil, fmt.Errorf("SendMessage failed: %w", UnknownError)
}
