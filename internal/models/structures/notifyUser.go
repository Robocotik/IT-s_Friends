package structures

import (
	"time"

	// "github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/go-co-op/gocron"
)

type NotifyUser struct {
	Id             int64
	NotifyInterval string
	NickName       string
}

func (nu NotifyUser) IsParamCustom() bool {
	return nu.NotifyInterval[len(nu.NotifyInterval)-1:] == "h" // custom param means that we need to calc the time for schedule
}

func (nu NotifyUser) NotifyCustom() {

}
func (nu NotifyUser) NotifyNoCustom() {

}

func (nu NotifyUser) WriteMessage() {
	// output.WriteMessage()
}

func (nu NotifyUser) Notify() {
	sh := gocron.NewScheduler(time.UTC)
	if nu.IsParamCustom() {
		sh.Every(1).At(nu.NotifyInterval).Do(func() {
			nu.NotifyCustom()
		})
	} else {
		sh.Every(1).At(nu.NotifyInterval).Do(func() {
			nu.NotifyNoCustom()
		})
	}
	sh.StartAsync()

}
