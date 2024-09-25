package keyboard

import (
	"Friends/src/assets"
	// "Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	// "fmt"
)

func CreateKeyboardFaculty() *telego.ReplyKeyboardMarkup {
	items := make([]telego.KeyboardButton, len(assets.Fakultets))
	for i := 0; i < len(assets.Fakultets); i++ {
		items[i] = tu.KeyboardButton(assets.Fakultets[i])
	}
	rowsCount := (len(assets.Fakultets) + 4) / 5
	items_rows := make([][]telego.KeyboardButton, rowsCount)

	for i := 0; i < rowsCount; i++ {
		start := i * 5
		end := start + 5
		if end > len(assets.Fakultets) {
			end = len(assets.Fakultets)
		}
		items_rows[i] = make([]telego.KeyboardButton, 0, end-start) // Создаем срез с нужной длиной
		for j := start; j < end; j++ {
			items_rows[i] = append(items_rows[i], tu.KeyboardButton(assets.Fakultets[j]))
		}
	}

	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите факультет").WithOneTimeKeyboard()
}
