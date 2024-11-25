package notify

import (
	"Friends/src/components/structures"
	// "Friends/src/utils"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mymmrac/telego"
)

func ParseUsers(conn *pgx.Conn, bot *telego.Bot, msg_id int64) []structures.NotifyUser {
	var result []structures.NotifyUser
	rows, err := conn.Query(context.Background(), "SELECT id, notify_interval FROM users")
	if err != nil {
		// utils.RiseError(bot, msg_id, err)
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var frequence time.Duration
		var id int64
		err = rows.Scan(&id, &frequence)
		fmt.Sprintf("ID: "+string(id)+"   Frequency: "+frequence.String() + "\n")
		// utils.WriteMessage(bot, msg, "ID: "+string(id)+"   Frequency: "+frequence.String() + "\n")
		if err != nil {
			fmt.Println(err)
			break
		}
		result = append(result, structures.NotifyUser{id, frequence})
		fmt.Println("ПОЛЬЗОВАТЕЛЬ,%s ВРЕМЯ %s", id, frequence)
	}
	return result
}
