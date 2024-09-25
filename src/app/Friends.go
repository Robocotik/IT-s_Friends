package main

import (
	"Friends/src/handlers"
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
			// Specify name (no validation)
			user.Faculty = msg.Text
			handle.HandleSelectCourse(bot, msg)
			user.State = StateAskFaculty
		case StateAskFaculty:
			// Specify age (validate that its positive number)
			// if err != nil || age == 0 {
			// 	_, _ = bot.SendMessage(tu.Message(
			// 		msg.Chat.ChatID(),
			// 		fmt.Sprintf("Invalid age, please try again"),
			// 	))
			// No state change
			// } else {
			user.Course = msg.Text
			handle.HandleSelectFaculty(bot, msg)
			user.State = StateAskCathedra
			// }
		case StateAskCathedra:
			// Specify email (validate that its valid email address)
			// var addr *mail.Address
			// addr, err = mail.ParseAddress(msg.Text)
			// if err != nil {
			// 	_, _ = bot.SendMessage(tu.Message(
			// 		msg.Chat.ChatID(),
			// 		fmt.Sprintf("Invalid email, please try again"),
			// 	))
			// No state change

			user.Faculty = msg.Text
			handle.HandleSelectCathedra(bot, msg)
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

	// bh.Handle(
	// 	handle.HandleSelectCourse,
	// 	th.CommandEqual("curs"),
	// )

	// bh.Handle(
	// 	handle.HandleSelectFaculty,
	// 	th.CommandEqual("fac"),
	// )

	// bh.Handle(
	// 	handle.HandleSelectCathedra,
	// 	th.CommandEqual("cathedra"),
	// )

	// bh.Handle(
	// 	handle.HandleStartCommand,
	// 	th.CommandEqual("start"),
	// )

	// bh.Handle(func(bot *telego.Bot, update telego.Update) {
	// 	// Send message
	// 	_, _ = bot.SendMessage(tu.Message(
	// 		tu.ID(update.Message.Chat.ID),
	// 		"Unknown command, use /start",
	// 	))
	// }, th.AnyCommand())

	bh.Start()
}
