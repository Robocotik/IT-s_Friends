package structures

type State uint

type User struct {
	State     State
	Filial    string
	Faculty   string
	Course    string
	Cathedra  string
	Group     string
	Favourite []string
}
