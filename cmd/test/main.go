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
	testSendMessageWithButtons()
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

func testSendMessageWithButtons() {
	ikm := telegram.CreateButtons(
		2,
		telegram.CreateCallbackButton("a button 1", "button 1 pressed"),
		telegram.CreateCallbackButton("a button 2", "button 2 pressed"),
		telegram.CreateCallbackButton("a button 3", "button 3 pressed"),
		telegram.CreateCallbackButton("a button 4", "button 4 pressed"),
		telegram.CreateLinkButton("justin.ooo", "https://justin.ooo/"),
	)

	var markup telegram.IMap
	mapstructure.Decode(ikm, &markup)

	message, err := client.SendMessage(JustinChatID, "test", &markup)
	if err != nil {
		utils.Exitf(0, "testSendMessage failed: %s", err)
	}

	utils.PrettyPrint(message)
}
