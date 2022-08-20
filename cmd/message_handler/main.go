package main

import (
	"time"

	"github.com/ooojustin/telego/pkg/telegram"
	"github.com/ooojustin/telego/pkg/utils"
)

func main() {
	cfg, ok := utils.GetConfig()
	if !ok {
		utils.Exitf(0, "Failed to load config.")
	}

	client := telegram.NewTelegramClient(cfg.TelegramToken)

	client.StartUpdateHandler(10*time.Second, []string{"message"})
}
