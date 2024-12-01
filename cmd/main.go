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
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

var (
	sessions      = make(map[int64]*structures.User)
	sessionsMutex sync.Mutex
)

func initEnv() {
	err := godotenv.Load("../configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initSessions(conn *pgx.Conn) {
	users, _ := database.GetAllIds(conn)
	for _, userId := range users {
		sessions[userId] = &structures.User{
			Id:     userId,
			Exists: true,
		}
	}
}

func main() {
	initEnv()
	botToken := os.Getenv("TOKEN")
	bot, err := telego.NewBot(botToken)
	botUser, _ := bot.GetMe()
	updates, _ := bot.UpdatesViaLongPolling(nil)
	conn, err := database.InitBD()
	defer conn.Close(context.Background())
	initSessions(conn)
	bh, _ := th.NewBotHandler(bot, updates)
	go notify.CronMain(conn, bot, botUser.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer bh.Stop()
	defer bot.StopLongPolling()
	bh.HandleMessage(func(bot *telego.Bot, msg telego.Message) {
		userID := msg.From.ID
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
