package telegram

import (
	"encoding/json"
	"errors"
)

// https://core.telegram.org/bots/api#user
type User struct {
	ID                    int    `json:"id"`
	IsBot                 bool   `json:"is_bot"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Username              string `json:"username"`
	LanguageCode          string `json:"language_code"`
	IsPremium             bool   `json:"is_premium"`
	AddedToAttachmentMenu bool   `json:"added_to_attachment_menu"`
}

type BotUser struct {
	User
	CanJoinGroups           bool `json:"can_join_groups"`
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool `json:"supports_inline_queries"`
}

type GetMeResponse struct {
	Ok          bool    `json:"ok"`
	Result      BotUser `mapstructure:"result"`
	Description string  `json:"description"`
}

func (tc *TelegramClient) GetMe() (*BotUser, error) {
	resp, err := tc.SendRequest(GET, "getMe", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data GetMeResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Ok {
		return &data.Result, nil
	} else if len(data.Description) > 0 {
		return nil, errors.New(data.Description)
	}

	return nil, UnknownError
}
