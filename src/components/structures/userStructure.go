package structures

type State uint

type User struct {
	State     State
	Id        int64
	Friend    AskedFriend
}
