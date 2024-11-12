package logic

import (
	"Friends/src/components/structures"
	"encoding/json"
	"fmt"
	"github.com/mymmrac/telego"

	"io/ioutil"
	"net/http"
)

func DoRequest(bot *telego.Bot, msg telego.Message, uid string) structures.Final_timetable {
	resp, err := http.Get("https://lks.bmstu.ru/lks-back/api/v1/schedules/groups/" + uid + "/public")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return structures.Final_timetable{} 
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var result structures.Final_timetable
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return structures.Final_timetable{} 
	}

	fmt.Println("RESP: ", result)
	return result
}
