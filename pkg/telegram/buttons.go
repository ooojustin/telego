package telegram

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
