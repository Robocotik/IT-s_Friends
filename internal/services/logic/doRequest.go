package logic

import (
	"encoding/json"
	"fmt"
	"github.com/Robocotik/IT-s_Friends/internal/models/entities"
	"github.com/mymmrac/telego"

	"io"
	"net/http"
)

func DoRequest(bot *telego.Bot, msg telego.Message, uid string) entities.Final_timetable {
	req := "https://lks.bmstu.ru/lks-back/api/v1/schedules/groups/" + uid + "/public"
	fmt.Println("ОТПРАВЛЯЮ ЗАПРОС НА ", req)
	resp, err := http.Get(req)
	if err != nil {
		fmt.Println("ERROR in getting schedule: ", err)
		return entities.Final_timetable{}
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR in readAll schedule: ", err)
		return entities.Final_timetable{}
	}
	var result entities.Final_timetable
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("ERROR in unmarshalling schedule: ", err)
		return entities.Final_timetable{}
	}

	return result
}
