package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SendPhotoResponse struct {
	TelegramResponse
	Result Message `json:"result"`
}

func (tc *TelegramClient) SendPhoto(chat int, photo string, caption string, replyMarkup *IMap) (*Message, error) {
	request := IMap{
		"chat_id": chat,
		"photo":   photo,
	}

	if len(caption) > 0 {
		request["caption"] = caption
	}

	if replyMarkup != nil {
		request["reply_markup"] = *replyMarkup
	}

	resp, err := tc.SendRequest(POST, "sendPhoto", &request)
	if err != nil {
		return nil, fmt.Errorf("SendPhoto failed: %w", err)
	}

	defer resp.Body.Close()

	var data SendPhotoResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("SendPhoto failed: %w", err)
	}

	if data.Ok {
		return &data.Result, nil
	} else if len(data.Description) > 0 {
		err = errors.New(data.Description)
		return nil, fmt.Errorf("SendPhoto failed: %w", err)
	}

	return nil, fmt.Errorf("SendPhoto failed: %w", UnknownError)
}
