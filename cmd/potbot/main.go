package main

import (
	"fmt"
	"os"

	"github.com/ooojustin/potbot/pkg/telegram"
	"github.com/ooojustin/potbot/pkg/utils"
)

func main() {
	cfg, ok := utils.GetConfig()
	if !ok {
		exitf(0, "Failed to load config.")
	}

	client := telegram.NewTelegramClient(cfg.TelegramToken)
	me, err := client.GetMe()
	if err != nil {
		exitf(0, "Failed to retrieve bot user info: %s", err)
	}

	fmt.Print(me)
}

func exitf(code int, format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(code)
}
