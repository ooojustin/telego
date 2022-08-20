package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
