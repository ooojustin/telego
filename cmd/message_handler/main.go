package main

import (
	"time"

	"github.com/ooojustin/telego/pkg/telegram"
	"github.com/ooojustin/telego/pkg/utils"
)

var (
	client *telegram.TelegramClient
)

func main() {
	cfg, ok := utils.GetConfig()
	if !ok {
		utils.Exitf(0, "Failed to load config.")
	}

	client = telegram.NewTelegramClient(cfg.TelegramToken)

	client.RegisterCommandHandler("ping", pingHandler, "ping command")
	client.RegisterCallbackQueryHandler(`button (\d+) pressed`, btnPressHandler)

	interval := time.Second * 10
	updateTypes := []string{"message", "callback_query"}

	client.StartUpdateHandler(interval, updateTypes)
}
