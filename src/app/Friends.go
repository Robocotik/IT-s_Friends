package main

import (
	"Friends/src/components/structures"
	"Friends/src/logic"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"

	th "github.com/mymmrac/telego/telegohandler"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	botToken := os.Getenv("TOKEN")
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	botUser, _ := bot.GetMe()
	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	var user structures.User

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Bot User: %+v\n", botUser)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.HandleMessage(func(bot *telego.Bot, msg telego.Message) {
		logic.DoSwitch(&user, bot, msg)
	})
	bh.Start()
}
