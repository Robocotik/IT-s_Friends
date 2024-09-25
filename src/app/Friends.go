package main

import (
	"Friends/src/handlers/handleSelectCourse"
	"Friends/src/handlers/handleSelectFaculty"
	"Friends/src/handlers/handleStart"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
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

type State uint

const (
	StateDefault State = iota
	StateAskCourse
	StateAskFaculty
	StateAskCathedra
	StateConfirm
	StateSearch
)

type User struct {
	State    State
	Faculty  string
	Course   uint
	Cathedra string
}

func main() {
	botToken := os.Getenv("TOKEN")
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	botUser, _ := bot.GetMe()
	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Bot User: %+v\n", botUser)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(
		handleCourse.HandleSelectCourse,
		th.CommandEqual("curs"),
	)

	bh.Handle(
		handleFaculty.HandleSelectFaculty,
		th.CommandEqual("fac"),
	)

	bh.Handle(
		handleStart.HandleStartCommand,
		th.CommandEqual("start"),
	)

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Unknown command, use /start",
		))
	}, th.AnyCommand())

	bh.Start()
}
