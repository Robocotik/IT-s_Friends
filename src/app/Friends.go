package main

import (
	"Friends/src/components/keyboard"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"os"
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
	keyboard := keyboard.CreateKeyboard()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	botUser, _ := bot.GetMe()
	fmt.Printf("Bot User: %+v\n", botUser)

	msg := tu.Message(
		tu.ID(123),
		"Hello World",
	).WithReplyMarkup(keyboard).WithProtectContent() // Multiple `with` method
	bot.SendMessage(msg)
}
