package main

import (
	"Friends/src/assets"
	"Friends/src/handlers"
	"Friends/src/messages"
	"Friends/src/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
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
	Course   string
	Cathedra string
}

func main() {
	botToken := os.Getenv("TOKEN")
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	botUser, _ := bot.GetMe()
	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	var user User

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Bot User: %+v\n", botUser)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.HandleMessage(func(bot *telego.Bot, msg telego.Message) {
		switch user.State {

		case StateDefault:
			handle.HandleStart(bot, msg)
			user.State = StateAskCourse

		case StateAskCourse:
			handle.HandleSelectCourse(bot, msg)
			if msg.Text != messages.Start {
				user.Faculty = utils.ParseString(bot, msg, "курс", assets.Fakultets[:])
				user.State = StateAskFaculty
			}
		case StateAskFaculty:
			handle.HandleSelectFaculty(bot, msg)
			user.Course = utils.ParseString(bot, msg, "факультет", assets.Courses[:])
			user.State = StateAskCathedra

		case StateAskCathedra:
			handle.HandleSelectCathedra(bot, msg)
			user.Cathedra = utils.ParseString(bot, msg, "кафедра", assets.Cathedras[:])
			user.State = StateConfirm

		case StateConfirm:
			_, _ = bot.SendMessage(tu.Message(
				msg.Chat.ChatID(),
				fmt.Sprintf(
					"Выходит твой друг учится на  %s, на  %s%s, Верно? (Y/N)",
					user.Course, user.Faculty, user.Cathedra,
				),
			))
			if msg.Text == "Y" {
				_, _ = bot.SendMessage(tu.Message(
					msg.Chat.ChatID(),
					fmt.Sprintf("Thanks for your data!"),
				))
				user.State = StateDefault
				// Do some logic
			} else {
				_, _ = bot.SendMessage(tu.Message(
					msg.Chat.ChatID(),
					fmt.Sprintf("Ok, let's start again"),
				))

				user.State = StateDefault
			}
		default:
			panic("unknown state")
		}
	})
	bh.Start()
}
