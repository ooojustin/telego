package main

import (
	"github.com/mitchellh/mapstructure"
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
	// testSendMessage()
	testSendMessageWithButton()
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
	message, err := client.SendMessage(JustinChatID, "test", nil)
	if err != nil {
		utils.Exitf(0, "testSendMessage failed: %s", err)
	}

	utils.PrettyPrint(message)
}

func testSendMessageWithButton() {
	ikm := telegram.InlineKeyboardMarkup{
		Keyboard: [][]telegram.InlineKeyboardButton{
			{
				telegram.CreateCallbackButton("a button", "button pressed"),
			},
		},
	}

	var markup telegram.IMap
	mapstructure.Decode(ikm, &markup)

	message, err := client.SendMessage(JustinChatID, "test", &markup)
	if err != nil {
		utils.Exitf(0, "testSendMessage failed: %s", err)
	}

	utils.PrettyPrint(message)
}
