package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Robocotik/IT-s_Friends/internal/database/postgres"
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
	users, err := postgres.GetAllIds(conn)
	if (err != nil){
		fmt.Println("Ошибка при получении всех пользователей ", err)
	}
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
	botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println("Ошибка при получении бота " + err.Error())
	}
	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		fmt.Println("Ошибка при update " + err.Error())
	}
	ctxMain, mainCancel := context.WithCancel(context.Background())
	defer mainCancel()
	ctxTimer, cancel := context.WithTimeout(ctxMain, 10*time.Second)
	defer cancel()
	conn, err := postgres.InitBD()
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
			postgres.AddUserId(bot, msg, conn, msg.Chat.ChatID().ID, msg.From.Username)
		}
		sessionsMutex.Unlock()

		go func() {
			logic.DoSwitch(ctxMain, conn, user, &user.Friend, bot, msg)
		}()
	})
	bh.Start()
	for {
		select {
		case <-ctxTimer.Done():
			fmt.Println("Ошибка при установки соединения с базой данных")
		default:
			fmt.Println("Соединение с базой данных установлено")
		}

	}
}
