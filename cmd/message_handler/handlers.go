package main

import (
	"fmt"
	"strconv"

	"github.com/ooojustin/telego/pkg/telegram"
)

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
	fmt.Println(msg)
	return nil
}
