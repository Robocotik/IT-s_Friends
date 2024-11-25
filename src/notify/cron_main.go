package notify

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func CronMain(conn *pgx.Conn, bot *telego.Bot, msg_id int64) {
    // Создаём новый планировщик
    s := gocron.NewScheduler(time.UTC)
    // Определяем задачу
    task := ParseUsers(conn, bot, msg_id)
    // Запускаем задачу каждые 5 секунд
    s.Every(5).Minutes().Do(task)
    // Запускаем планировщик в отдельной горутине
    go func() {
        s.StartAsync()
    }()
    // Держим приложение в рабочем состоянии
    select {}
}