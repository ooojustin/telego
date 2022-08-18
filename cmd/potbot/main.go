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
}

func exitf(code int, format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(code)
}
