package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Robocotik/IT-s_Friends/internal/database"
	"github.com/Robocotik/IT-s_Friends/internal/database/postgres"
	"github.com/Robocotik/IT-s_Friends/internal/models/structures"
	"github.com/Robocotik/IT-s_Friends/internal/services/logic"
	"github.com/Robocotik/IT-s_Friends/internal/services/notify"
	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func initEnv() {
	err := godotenv.Load("../configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initEnv()

	sessions := make(map[int64]*structures.User)
	var sessionsMutex sync.Mutex
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
	var psql postgres.Postgres
	psql.Conn, err = database.NewDatabase("POSTGRES_DRIVER", "POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_TABLE")
	defer psql.Conn.Close(context.Background())
	psql.InitSessions(&sessions)
	bh, _ := th.NewBotHandler(bot, updates)
	go notify.CronMain(psql.Conn, bot, botUser.ID)
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
			psql.AddUserId(bot, msg, msg.Chat.ChatID().ID, msg.From.Username)
		}
		sessionsMutex.Unlock()

		go func() {
			logic.DoSwitch(ctxMain, psql, user, &user.Friend, bot, msg)
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
