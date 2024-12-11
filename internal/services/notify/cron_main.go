package notify

import (
	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func CronMain(conn *pgx.Conn, bot *telego.Bot, chatID int64) {
	users := ParseUsers(conn, bot, chatID)
	for _, user := range users {
        user.Notify()
	}
}

// s := gocron.NewScheduler(time.UTC)
// 	task := ParseUsers(conn, bot, chatID)
// Запускаем задачу каждые 5 секунд
// s.Every(5).Minutes().Do(task)
// Запускаем планировщик в отдельной горутине
