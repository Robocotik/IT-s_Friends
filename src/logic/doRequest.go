package logic

import (
	"fmt"
	"io"
	"net/http"
	// "encoding/json"
)

func DoRequest(uid string) {
	resp, err := http.Get("https://lks.bmstu.ru/lks-back/api/v1/schedules/groups/" + uid + "/public")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer resp.Body.Close()
	// Data json
	body, err := io.ReadAll(resp.Body)
	fmt.Println("RESP: ", body)
}
