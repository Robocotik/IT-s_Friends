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
    
	err_del := dropStaticDataBD(ctxx, conn)
	if err_del != nil {
        fmt.Println("Ошибка при удалении старой информцации БД:", err_del)
        return err_del
    }
    err := ParseAllSchdule(ctxx, conn, bot, chatID)
    if err != nil {
        fmt.Println("Ошибка при обновлении БД:", err)
        return err
    }
    
    fmt.Println("База данных успешно обновлена")
    return nil
}