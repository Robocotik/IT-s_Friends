package main

import (
	"github.com/Robocotik/IT-s_Friends/src/components/structures"
	"github.com/Robocotik/IT-s_Friends/src/logic"
	"github.com/Robocotik/IT-s_Friends/src/notify"
	"github.com/Robocotik/IT-s_Friends/src/utils/server"
	"context"
	"fmt"
	"log"
	"os"
	"sync"

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
	var UserSessions = make(map[int64]*structures.User) // Используйте int64 для ID пользователя
	var mu sync.Mutex                                   // Мьютекс для синхронизации доступа к UserSessions
	botToken := os.Getenv("TOKEN")
	bot, err := telego.NewBot(botToken)
	botUser, _ := bot.GetMe()
	updates, _ := bot.UpdatesViaLongPolling(nil)
	conn, err := server.InitBD()
	defer conn.Close(context.Background())
	bh, _ := th.NewBotHandler(bot, updates)
	go notify.CronMain(conn, bot, botUser.ID )
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Printf("Bot User: %+v\n", botUser)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.HandleMessage(func(bot *telego.Bot, msg telego.Message) {
		mu.Lock() // Блокируем доступ к UserSessions
		defer mu.Unlock()

		userID := msg.From.ID
		// Проверяем, есть ли уже сессия для этого пользователя
		user, exists := UserSessions[userID]
		if !exists {
			// Если сессии нет, создаем новую
			user = &structures.User{
				Id:    userID,
				State: structures.StateStart,
				// Заполните другие поля структуры, если необходимо
			}
			UserSessions[userID] = user
			server.AddUserId(bot, msg, conn, msg.Chat.ChatID().ID, msg.From.Username)
		}
		logic.DoSwitch(conn, UserSessions[userID], &UserSessions[userID].Friend, bot, msg, exists)
	})
	bh.Start()
}
