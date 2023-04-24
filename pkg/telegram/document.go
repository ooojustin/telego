package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
)

// https://core.telegram.org/bots/api#sending-files
// https://core.telegram.org/bots/api#senddocument
func (tc *TelegramClient) SendDocument(chat int, document string, caption string, replyMarkup *IMap) (*Message, error) {
	request := IMap{"chat_id": chat}

	request["document"] = document // file url
	request["caption"] = caption

	if replyMarkup != nil {
		request["reply_markup"] = *replyMarkup
	}

	resp, err := tc.SendRequest(POST, "sendDocument", &request)
	if err != nil {
		return nil, fmt.Errorf("SendDocument failed: %w", err)
	}

	defer resp.Body.Close()

	var data SendMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("SendDocument failed: %w", err)
	}

	if data.Ok {
		return &data.Result, nil
	} else if len(data.Description) > 0 {
		err = errors.New(data.Description)
		return nil, fmt.Errorf("SendDocument failed: %w", err)
	}

	return nil, fmt.Errorf("SendDocument failed: %w", UnknownError)
}
