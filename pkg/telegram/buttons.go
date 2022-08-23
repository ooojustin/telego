package telegram

import (
	"math"
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
