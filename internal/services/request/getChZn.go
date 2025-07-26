package request

// import (
// 	"github.com/Robocotik/IT-s_Friends/internal/models/entities"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func GetChZn() string {
// 	resp, err := http.Get("https://lks.bmstu.ru/lks-back/api/v1/schedules/current")
// 	if err != nil {
// 		fmt.Println("ERROR in getting chzn: ", err)
// 		return "-1"
// 	}

// 	defer resp.Body.Close()
// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("ERROR in readAll chzn: ", err)
// 		return "-1"
// 	}
// 	var result entities.Final_chzn
// 	err = json.Unmarshal(data, &result)
// 	if err != nil {
// 		fmt.Println("ERROR in unmaarshalling chzn: ", err)
// 		return "-1"
// 	}

// 	fmt.Println("CHZN: ", result.Data.WeekName)
// 	return result.Data.WeekName
