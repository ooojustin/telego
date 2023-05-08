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
	BotCommandScopeAllPrivateChats Scope = Scope{Type: "all_private_chats"}
)

type Scope struct {
	Type string `json:"type"`
}

type BotCommand struct {
	Command     string `json:"command"`     // Max 32 chars
	Description string `json:"description"` // Max 256 chars
	IsCategory  bool   `json:"-"`
}

type SetMyCommandsResponse struct {
	TelegramResponse
	Result bool `json:"result"`
}

type GetMyCommandsResponse struct {
	TelegramResponse
	Result []BotCommand `json:"result"`
}

type DeleteMyCommandsResponse struct {
	TelegramResponse
	Result bool `json:"result"`
}

// https://core.telegram.org/bots/api#setmycommands
func (tc *TelegramClient) SetMyCommands(commands []BotCommand, scope *Scope, languageCode string) (bool, error) {
	if len(languageCode) == 0 {
		languageCode = DefaultLanguageCode
	}

	request := IMap{
		"commands":      commands,
		"language_code": languageCode,
	}

	if scope == nil {
		request["scope"] = &BotCommandScopeAllPrivateChats
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
func (tc *TelegramClient) GetMyCommands(scope *Scope, languageCode string) (*[]BotCommand, error) {
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
func (tc *TelegramClient) DeleteMyCommands(scope *Scope, languageCode string) (bool, error) {
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

func (tc *TelegramClient) RegisterCommandHandler(command string, handler UpdateHandler, description string) {
	tc.CommandHandlers[command] = handler
	if len(description) > 0 {
		tc.CommandDescriptions[command] = description
	} else {
		tc.CommandDescriptions[command] = "[no description provided]"
	}
}

func (tc *TelegramClient) RegisterCallbackQueryHandler(dataPattern string, handler CallbackQueryHandler) {
	tc.CallbackQueryHandlers[dataPattern] = handler
}

func (tc *TelegramClient) GetBotCommands() []BotCommand {
	var commands []BotCommand
	for command, description := range tc.CommandDescriptions {
		commands = append(commands, BotCommand{
			Command:     command,
			Description: description,
		})
	}
	return commands
}
