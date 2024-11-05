package main

import (
	"Friends/src/assets"
	"Friends/src/handlers"
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
	StateStart State = iota
	StateDefault
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

		case StateStart:

			msg.Text = ""
			handle.HandleStart(bot, msg)
			user.State = StateDefault
			msg.Text = ""

		case StateDefault:

			user.Course = utils.ParseString(bot, msg, "Погнали", []string{"Погнали"})
			handle.HandleSelectCourse(bot, msg)
			user.State = StateAskCourse

		case StateAskCourse:
			
			user.Course = utils.ParseString(bot, msg, "курс", assets.Courses[:])
			handle.HandleSelectFaculty(bot, msg)
			user.State = StateAskFaculty

		case StateAskFaculty:

			user.Faculty = utils.ParseString(bot, msg, "факультет", assets.Fakultets[:])
			handle.HandleSelectCathedra(bot, msg)
			user.State = StateAskCathedra

		case StateAskCathedra:

			user.Cathedra = utils.ParseString(bot, msg, "кафедра", assets.Cathedras[:])
			// handle.HandleSelectCathedra(bot, msg)
			user.State = StateConfirm
			_, _ = bot.SendMessage(tu.Message(
				msg.Chat.ChatID(),
				fmt.Sprintf(
					"Выходит твой друг учится на  %s, на  %s%s, Верно? (Y/N)",
					user.Course, user.Faculty, user.Cathedra,
				),
			))
		case StateConfirm:

			if msg.Text == "Y" {
				_, _ = bot.SendMessage(tu.Message(
					msg.Chat.ChatID(),
					fmt.Sprintf("Thanks for your data!"),
				))
				user.State = StateStart
				// Do some logic
			} else {
				_, _ = bot.SendMessage(tu.Message(
					msg.Chat.ChatID(),
					fmt.Sprintf("Ok, let's start again"),
				))

				user.State = StateStart
			}
		default:
			_, _ = bot.SendMessage(tu.Message(
				msg.Chat.ChatID(),
				fmt.Sprintf("Неизвестная команда"),
			))
			panic("unknown state")
		}
	})
	bh.Start()
}
