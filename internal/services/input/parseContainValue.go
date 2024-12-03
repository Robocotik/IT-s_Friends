package input

import (
	"errors"
	"strings"
	// "github.com/Robocotik/IT-s_Friends/internal/services/output"
	"github.com/mymmrac/telego"
)

func ParseContainString(bot *telego.Bot, msg telego.Message, possibleData []string) (string, error) {
	found := false
	for _, data := range possibleData {
		if strings.Contains(msg.Text, data) {
			found = true
			break
		}
	}
	if !found {
		// output.RiseError(bot, msg, err)
		return "", errors.New("not found")
	}
	return msg.Text, nil
}
