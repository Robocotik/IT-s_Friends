package utils

import (
	"Friends/src/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetChZn() string {
	resp, err := http.Get("https://lks.bmstu.ru/lks-back/api/v1/schedules/current")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return "-1"
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var result entities.Final_chzn
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return "-1"
	}

	fmt.Println("CHZN: ", result.Data.WeekName)
	return result.Data.WeekName
}
