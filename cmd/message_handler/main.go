package main

import (
	"fmt"
	"strconv"
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

	client.RegisterCommandHandler("ping", pingHandler)
	client.RegisterCallbackQueryHandler("button (\\d+) pressed", btnPressHandler)

	interval := time.Second * 10
	updateTypes := []string{"message", "callback_query"}

	client.StartUpdateHandler(interval, updateTypes)
}

func pingHandler(update telegram.Update) error {
	chatId := update.Message.Chat.ID
	client.SendMessage(chatId, "pong!", nil)
	fmt.Println("ping command handled!")
	return nil
}

func btnPressHandler(update telegram.Update, groups []string) error {
	btnNum, err := strconv.Atoi(groups[1])
	if err != nil {
		return nil
	}
	msg := fmt.Sprintf("button press handled: #%d", btnNum)
	chatId := update.CallbackQuery.Message.Chat.ID
	client.SendMessage(chatId, msg, nil)
	fmt.Println("button press handled!")
	return nil
}
