package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"

	"github.com/ooojustin/telego/pkg/utils"
)

type MenuButtonType string

const (
	MENU_BUTTON_COMMANDS MenuButtonType = "commands"
	MENU_BUTTON_DEFAULT  MenuButtonType = "default"
	MENU_BUTTON_WEB_APP  MenuButtonType = "web_app"
)

var (
	InvalidMenuButtonType error = errors.New("GetChatMenuButton failed: invalid MenuButtonType")
)

// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	Keyboard [][]InlineKeyboardButton `mapstructure:"inline_keyboard" json:"inline_keyboard"`
}

// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	Url          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"` // Up to 64 bytes
}

// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	Id      string  `json:"id"`
	From    User    `json:"from"`
	Message Message `json:"message"`
	ChatID  string  `json:"chat_instance"`
	Data    string  `json:"data"`
}

// https://core.telegram.org/bots/api#menubutton
type MenuButton struct {
	Type MenuButtonType `json:"type"`
}

// https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	Url string `json:"url,omitempty"`
}

// https://core.telegram.org/bots/api#menubuttonwebapp
type MenuButtonWebApp struct {
	MenuButton
	Text   string     `json:"text,omitempty"`
	WebApp WebAppInfo `json:"web_app,omitempty"`
}

type SetChatMenuButtonResponse struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

func CreateLinkButton(text string, url string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		Url:  url,
	}
}

func CreateCallbackButton(text string, data string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	}
}

func CreateButtons(perRow int, buttons ...InlineKeyboardButton) InlineKeyboardMarkup {
	btnCount := len(buttons)
	btnsPerRow := float64(btnCount) / float64(perRow)
	rows := int(math.Ceil(btnsPerRow))

	keyboard := [][]InlineKeyboardButton{}
	for idx := 0; idx < rows; idx++ {
		if btnCount >= perRow {
			// add row with the first 'perRow' buttons
			row := buttons[:perRow]
			keyboard = append(keyboard, row)

			// remove used buttons, update count
			buttons = buttons[perRow:]
			btnCount = len(buttons)
		} else {
			// add final row of buttons
			keyboard = append(keyboard, buttons)
		}
	}

	return InlineKeyboardMarkup{Keyboard: keyboard}
}

func (tc *TelegramClient) DeleteChatMenuButton(chatId *int) error {
	return _setChatMenuButton(tc, chatId, nil)
}

func (tc *TelegramClient) SetChatMenuButton(chatId *int, mbType MenuButtonType) error {
	menuButton := MenuButton{Type: mbType}
	return _setChatMenuButton(tc, chatId, &menuButton)
}

func (tc *TelegramClient) SetChatWebAppMenuButton(chatId *int, menuButton MenuButtonWebApp) error {
	menuButton.MenuButton = MenuButton{
		Type: MENU_BUTTON_WEB_APP,
	}
	return _setChatMenuButton(tc, chatId, &menuButton)
}

func (tc *TelegramClient) GetChatMenuButton(chatId *int) (IMap, MenuButtonType, error) {
	request := IMap{}

	if chatId != nil {
		request["chat_id"] = *chatId
	}

	resp, err := tc.SendRequest(POST, "getChatMenuButton", &request)
	if err != nil {
		return IMap{}, "", fmt.Errorf("GetChatMenuButton failed: %w", err)
	}

	defer resp.Body.Close()

	var data IMap
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return IMap{}, "", fmt.Errorf("GetChatMenuButton failed: %w", err)
	}

	if resultI, ok := data["result"]; ok {
		result := resultI.(map[string]interface{})
		if mbType, ok := result["type"]; ok {
			switch mbType.(string) {
			case string(MENU_BUTTON_COMMANDS):
				return data, MENU_BUTTON_COMMANDS, nil
			case string(MENU_BUTTON_DEFAULT):
				return data, MENU_BUTTON_DEFAULT, nil
			case string(MENU_BUTTON_WEB_APP):
				return data, MENU_BUTTON_WEB_APP, nil
			}
			return data, "", InvalidMenuButtonType
		}
	}

	return IMap{}, "", InvalidMenuButtonType
}

func _setChatMenuButton(tc *TelegramClient, chatId *int, menuButton interface{}) error {
	request := IMap{}

	if !utils.IsNil(menuButton) {
		request["menu_button"] = menuButton
	}

	if chatId != nil {
		request["chat_id"] = *chatId
	}

	resp, err := tc.SendRequest(POST, "setChatMenuButton", &request)
	if err != nil {
		return fmt.Errorf("SetChatMenuButton failed: %w", err)
	}

	defer resp.Body.Close()

	var data SetChatMenuButtonResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return fmt.Errorf("SetChatMenuButton failed: %w", err)
	}

	if data.Ok {
		return nil
	}

	return errors.New("SetChatMenuButton failed")
}
