package main

import (
	"fmt"
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

	client.RegisterCommandHandler("test", testCommandHandler)

	interval := time.Second * 10
	updateTypes := []string{"message", "callback_query"}

	client.StartUpdateHandler(interval, updateTypes)
}

func testCommandHandler(update telegram.Update) error {
	fmt.Println("test command executed!")
	return nil
}
