package main

import (
	"github.com/ooojustin/telego/pkg/telegram"
	"github.com/ooojustin/telego/pkg/utils"
)

func main() {
	cfg, ok := utils.GetConfig()
	if !ok {
		utils.Exitf(0, "Failed to load config.")
	}

	client := telegram.NewTelegramClient(cfg.TelegramToken)

	testGetMe(client)

	// testGetUpdates(client)
}

func testGetMe(client *telegram.TelegramClient) {
	me, err := client.GetMe()
	if err != nil {
		utils.Exitf(0, "testGetMe failed: %s", err)
	}

	utils.PrettyPrint(*me)
}

func testGetUpdates(client *telegram.TelegramClient) {
	updates, err := client.GetUpdates(0, []string{"message"})
	if err != nil {
		utils.Exitf(0, "testGetUpdates failed: %s", err)
	}

	utils.PrettyPrint(updates)
}
