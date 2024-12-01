package structures

import "time"

type NotifyUser struct {
	Id               int64
	Notify_interval time.Duration
}
