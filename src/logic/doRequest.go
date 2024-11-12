package logic

import (
	"Friends/src/components/structures"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DoRequest(uid string) structures.Final_timetable {
	resp, err := http.Get("https://lks.bmstu.ru/lks-back/api/v1/schedules/groups/" + uid + "/public")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var result structures.Final_timetable
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	// Data json
	fmt.Println("RESP: ", result)
	return result
}
