package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
)

// https://core.telegram.org/bots/api#user
type User struct {
	ID                    int    `json:"id" mapstructure:"id"`
	IsBot                 bool   `json:"is_bot" mapstructure:"is_bot"`
	FirstName             string `json:"first_name" mapstructure:"first_name"`
	LastName              string `json:"last_name" mapstructure:"last_name"`
	Username              string `json:"username" mapstructure:"username"`
	LanguageCode          string `json:"language_code" mapstructure:"language_code"`
	IsPremium             bool   `json:"is_premium" mapstructure:"is_premium"`
	AddedToAttachmentMenu bool   `json:"added_to_attachment_men" mapstructure:"added_to_attachment_menu"`
}

type BotUser struct {
	User
	CanJoinGroups           bool `json:"can_join_groups" mapstructure:"can_join_groups"`
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages" mapstructure:"can_read_all_group_messages"`
	SupportsInlineQueries   bool `json:"supports_inline_queries" mapstructure:"supports_inline_queries"`
}

type GetMeResponse struct {
	TelegramResponse
	Result BotUser `json:"result"`
}

// https://core.telegram.org/bots/api#getme
func (tc *TelegramClient) GetMe() (*BotUser, error) {
	resp, err := tc.SendRequest(GET, "getMe", nil)
	if err != nil {
		return nil, fmt.Errorf("GetMe failed: %w", err)
	}

	defer resp.Body.Close()

	var data GetMeResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("GetMe failed: %w", err)
	}

	if data.Ok {
		return &data.Result, nil
	} else if len(data.Description) > 0 {
		err = errors.New(data.Description)
		return nil, fmt.Errorf("GetMe failed: %w", err)
	}

	return nil, fmt.Errorf("GetMe failed: %w", UnknownError)
}
