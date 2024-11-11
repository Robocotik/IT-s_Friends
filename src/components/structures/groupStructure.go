package structures

type Group struct {
	Abbr       string `json:"abbr"`
	Uuid       string `json:"uuid"`
	Course     uint   `json:"course"`
	NodeType   string `json:"nodeType"`
	Semester   uint   `json:"semestr"`
	ParentUuid string `json:"parentUuid"`
}

type Course struct {
	Abbr       string  `json:"abbr"`
	Course     uint    `json:"course"`
	NodeType   string  `json:"nodeType"`
	ParentUuid string  `json:"parentUuid"`
	Children   []Group `json:"children"`
}

type Cathedra struct {
	Abbr       string   `json:"abbr"`
	Name       string   `json:"name"`
	Uuid       string   `json:"uuid"`
	NodeType   string   `json:"nodeType"`
	ParentUuid string   `json:"parentUuid"`
	Children   []Course `json:"children"`
}

type Faculty struct {
	Abbr     string     `json:"abbr"`
	Name     string     `json:"name"`
	Uuid     string     `json:"uuid"`
	Children []Cathedra `json:"children"` // Изменено на массив Group
}

type Fillial struct {
	Abbr     string    `json:"abbr"`
	Name     string    `json:"name"`
	Uuid     string    `json:"uuid"`
	Children []Faculty `json:"children"`
}

type IData struct {
	Abbr     string    `json:"abbr"`
	Name     string    `json:"name"`
	Uuid     string    `json:"uuid"`
	Children []Fillial `json:"children"`
}

type Final struct {
	Data IData `json:"data"`
}
