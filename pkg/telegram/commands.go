package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	DefaultLanguageCode = "en"
)

var (
	BotCommandScopeAllPrivateChats IMap = IMap{"type": "all_private_chats"}
)

type BotCommand struct {
	Command     string `json:"command"`     // Max 32 chars
	Description string `json:"description"` // Max 256 chars
}

type SetMyCommandsResponse struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

type GetMyCommandsResponse struct {
	Ok          bool         `json:"ok"`
	Result      []BotCommand `mapstructure:"result"`
	Description string       `json:"description"`
}

type DeleteMyCommandsResponse struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

// https://core.telegram.org/bots/api#setmycommands
func (tc *TelegramClient) SetMyCommands(commands []BotCommand, scope *IMap, languageCode string) (bool, error) {
	if scope == nil {
		scope = &BotCommandScopeAllPrivateChats
	}

	if len(languageCode) == 0 {
		languageCode = DefaultLanguageCode
	}

	request := IMap{
		"commands":      commands,
		"scope":         scope,
		"language_code": languageCode,
	}

	resp, err := tc.SendRequest(POST, "setMyCommands", &request)
	if err != nil {
		return false, fmt.Errorf("SetMyCommands failed: %w", err)
	}

	defer resp.Body.Close()

	var data SetMyCommandsResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return false, fmt.Errorf("SetMyCommands failed: %w", err)
	}

	if data.Ok {
		return data.Result, nil
	} else if len(data.Description) > 0 {
		err = errors.New(data.Description)
		return false, fmt.Errorf("SetMyCommands failed: %w", err)
	}

	return false, fmt.Errorf("SetMyCommands failed: %w", UnknownError)
}

// https://core.telegram.org/bots/api#getmycommands
func (tc *TelegramClient) GetMyCommands(scope *IMap, languageCode string) (*[]BotCommand, error) {
	if scope == nil {
		scope = &BotCommandScopeAllPrivateChats
	}

	if len(languageCode) == 0 {
		languageCode = DefaultLanguageCode
	}

	request := IMap{
		"scope":         scope,
		"language_code": languageCode,
	}

	resp, err := tc.SendRequest(POST, "getMyCommands", &request)
	if err != nil {
		return nil, fmt.Errorf("GetMyCommands failed: %w", err)
	}

	defer resp.Body.Close()

	var data GetMyCommandsResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("GetMyCommands failed: %w", err)
	}

	if data.Ok {
		return &data.Result, nil
	} else if len(data.Description) > 0 {
		err = errors.New(data.Description)
		return nil, fmt.Errorf("GetMyCommands failed: %w", err)
	}

	return nil, fmt.Errorf("GetMyCommands failed: %w", UnknownError)
}

// https://core.telegram.org/bots/api#deletemycommands
func (tc *TelegramClient) DeleteMyCommands(scope *IMap, languageCode string) (bool, error) {
	if scope == nil {
		scope = &BotCommandScopeAllPrivateChats
	}

	if len(languageCode) == 0 {
		languageCode = DefaultLanguageCode
	}

	request := IMap{
		"scope":         scope,
		"language_code": languageCode,
	}

	resp, err := tc.SendRequest(POST, "deleteMyCommands", &request)
	if err != nil {
		return false, fmt.Errorf("DeleteMyCommands failed: %w", err)
	}

	defer resp.Body.Close()

	var data DeleteMyCommandsResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return false, fmt.Errorf("DeleteMyCommands failed: %w", err)
	}

	if data.Ok {
		return data.Result, nil
	} else if len(data.Description) > 0 {
		err = errors.New(data.Description)
		return false, fmt.Errorf("DeleteMyCommands failed: %w", err)
	}

	return false, fmt.Errorf("DeleteMyCommands failed: %w", UnknownError)
}
