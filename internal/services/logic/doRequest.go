package logic

import (
	"github.com/Robocotik/IT-s_Friends/internal/models/web"
	"encoding/json"
	"fmt"
	"github.com/mymmrac/telego"

	"io/ioutil"
	"net/http"
)

func DoRequest(bot *telego.Bot, msg telego.Message, uid string) entities.Final_timetable {
	resp, err := http.Get("https://lks.bmstu.ru/lks-back/api/v1/schedules/groups/" + uid + "/public")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return entities.Final_timetable{} 
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var result entities.Final_timetable
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return entities.Final_timetable{} 
	}

	return result
}
