package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/Robocotik/IT-s_Friends/internal/database"
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/logic"
	"github.com/Robocotik/IT-s_Friends/internal/services/notify"
	"github.com/joho/godotenv"

	"github.com/mymmrac/telego"

	th "github.com/mymmrac/telego/telegohandler"
)

var (
	sessions      = make(map[int64]*structures.User)
	sessionsMutex sync.Mutex
)

func init() {
	err := godotenv.Load("../configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {
	fmt.Sprintf("Я запустился")
	botToken := os.Getenv("TOKEN")
	bot, err := telego.NewBot(botToken)
	botUser, _ := bot.GetMe()
	updates, _ := bot.UpdatesViaLongPolling(nil)
	conn, err := database.InitBD()
	defer conn.Close(context.Background())
	bh, _ := th.NewBotHandler(bot, updates)
	go notify.CronMain(conn, bot, botUser.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Printf("Bot User: %+v\n", botUser)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.HandleMessage(func(bot *telego.Bot, msg telego.Message) {
		userID := msg.From.ID

		// Блокируем доступ к UserSessions только для получения или создания сессии
		sessionsMutex.Lock()
		user, exists := sessions[userID]
		if !exists {
			user = &structures.User{
				Id:     userID,
				State:  structures.StateStart,
				Exists: false,
			}
			sessions[userID] = user
			database.AddUserId(bot, msg, conn, msg.Chat.ChatID().ID, msg.From.Username)
		}
		sessionsMutex.Unlock()

		go func() {
			logic.DoSwitch(conn, user, &user.Friend, bot, msg)
		}()
	})
	bh.Start()
}
