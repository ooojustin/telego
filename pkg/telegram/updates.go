package telegram

// https://core.telegram.org/bots/api#update
type Update struct {
	ID      int     `json:"update_id"`
	Message Message `mapstructure:"message"`
}
