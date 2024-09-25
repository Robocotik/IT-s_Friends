package keyboard

import (
	"Friends/src/assets/faculty"
	// "Friends/src/messages"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	// "fmt"
)

func CreateKeyboardFaculty() *telego.ReplyKeyboardMarkup {
	items := make([]telego.KeyboardButton, len(faculty.Fakultets))
	for i := 0; i < len(faculty.Fakultets); i++ {
		items[i] = tu.KeyboardButton(faculty.Fakultets[i])
	}
	rowsCount := (len(faculty.Fakultets) + 4) / 5
	items_rows := make([][]telego.KeyboardButton, rowsCount)

	for i := 0; i < rowsCount; i++ {
		start := i * 5
		end := start + 5
		if end > len(faculty.Fakultets) {
			end = len(faculty.Fakultets)
		}
		items_rows[i] = make([]telego.KeyboardButton, 0, end-start) // Создаем срез с нужной длиной
		for j := start; j < end; j++ {
			items_rows[i] = append(items_rows[i], tu.KeyboardButton(faculty.Fakultets[j]))
		}
	}

	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите курс")
}
