package structures

type State uint

type User struct {
	State     State
	Favourite []Fav
	Id        int64
	Friend    AskedFriend
}
