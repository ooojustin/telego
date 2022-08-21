package main

import (
	"github.com/ooojustin/telego/pkg/telegram"
	"github.com/ooojustin/telego/pkg/utils"
)

const (
	JustinChatID int = 391089352
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

	// testGetMe()
	// testGetUpdates()
	testSendMessage()
}

func testGetMe() {
	me, err := client.GetMe()
	if err != nil {
		utils.Exitf(0, "testGetMe failed: %s", err)
	}

	utils.PrettyPrint(*me)
}

func testGetUpdates() {
	updates, err := client.GetUpdates(0, []string{"message"})
	if err != nil {
		utils.Exitf(0, "testGetUpdates failed: %s", err)
	}

	utils.PrettyPrint(updates)
}

func testSendMessage() {
	message, err := client.SendMessage(JustinChatID, "test")
	if err != nil {
		utils.Exitf(0, "testSendMessage failed: %s", err)
	}

	utils.PrettyPrint(message)
}
