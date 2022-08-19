package telegram

// https://core.telegram.org/bots/api#message
type Message struct {
	ID         int    `json:"message_id"`
	From       User   `json:"from"`
	SenderChat Chat   `json:"sender_chat"`
	Date       int    `json:"date"`
	Chat       Chat   `json:"chat"`
	Text       string `json:"text"`
}
