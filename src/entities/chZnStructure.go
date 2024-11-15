package entities

type IData_chzn struct {
	WeekName      string `json:"weekName"`
	WeekNumber    int    `json:"weekNumber"`
	WeekShortName string `json:"weekShortName"`
}

type Final_chzn struct {
	Data IData_chzn  `json:"data"`
	Date string `json:"date"`
}

