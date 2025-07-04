package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func dropStaticDataBD(ctx context.Context, conn *pgx.Conn) error {
	var err error
	dropTables := []string{
		"fillials",
		"cathedras",
		"faculties",
		"groups",
		"courses",
		"schedule",
	}
	for _, table := range dropTables {
		_, err = conn.Exec(ctx, fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table))
		if err != nil {
			fmt.Println("Ошибка при удалении данных таблицы:", err)
			break
		}
		fmt.Println("Успешно удалил данные таблицы:", table)
	}

	return err
}