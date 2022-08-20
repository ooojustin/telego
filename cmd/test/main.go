package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ooojustin/telego/pkg/telegram"
	"github.com/ooojustin/telego/pkg/utils"
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

	meBytes, err := json.MarshalIndent(*me, "", "	")
	if err == nil {
		meStr := string(meBytes)
		fmt.Println(meStr)
	}

	updates, _ := client.GetUpdates([]string{})
	fmt.Println(updates)
}

func exitf(code int, format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(code)
}
