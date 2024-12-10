package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func UpdateBd(ctx context.Context, conn *pgx.Conn, bot *telego.Bot, chatID int64) error {
	ctxx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	err := ParseAllSchdule(ctxx, conn, bot, chatID)

	for {
		select {
		case <-ctxx.Done():
			fmt.Println("Превышено время ожидания при обращении к бд")
			return err
		default:
			fmt.Println("База данных успешно обновлена")
		}
	}
	
}
