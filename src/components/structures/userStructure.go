package structures

import "time"

type State uint

type User struct {
	State          State
	StateFilling   State
	Exists         bool
	Id             int64
	Friend         AskedFriend
	Identity       Identity
	NotifyInterval time.Duration
}
