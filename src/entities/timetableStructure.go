package entities

type IDiscipline struct {
	Abbr      string `json:"abbr"`
	ActType   string `json:"actType"`
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
}

type IAudience struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

type ITeacher struct {
	Uuid       string `json:"uuid"`
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
}

type IGroup struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

type IDay struct {
	Day        int      `json:"day"`
	Time       int      `json:"time"`
	Week       string      `json:"week"`
	Groups     []IGroup    `json:"groups"`
	Stream     string      `json:"stream"`
	EndTime    string      `json:"endTime"`
	Teachers   []ITeacher  `json:"teachers"`
	Audiences  []IAudience `json:"audiences"`
	StartTime  string      `json:"startTime"`
	Discipline IDiscipline `json:"discipline"`
	Permission string      `json:"permission"`
}

type IData_timetable struct {
	Type     string `json:"type"`
	Uuid     string `json:"uuid"`
	Title    string `json:"title"`
	Schedule []IDay  `json:"schedule"`
}

type Final_timetable struct {
	Data IData_timetable `json:"data"`
	Date string          `json:"date"`
}
