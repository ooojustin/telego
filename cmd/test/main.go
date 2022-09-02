package main

import (
	"fmt"

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
	// testSendMessageWithButtons()
	// testSetWebhook()
	testGetWebhookInfo()
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

func testSetWebhook() {
	updateTypes := []string{"message", "callback_query"}
	url := "https://api.sellegram.net/telegram/webhook"
	token := "512a0f59-81b8-4648-9ab6-c312765b31c1"
	if err := client.SetWebhook(url, updateTypes, token); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("setWebhook successful.")
	}
}

func testGetWebhookInfo() {
	if info, err := client.GetWebhookInfo(); err != nil {
		fmt.Println(err)
	} else {
		prettyInfoStr, _ := utils.GetPrettyJSON(info)
		fmt.Printf("GetWebhookInfo: %s\n", prettyInfoStr)
	}
}
