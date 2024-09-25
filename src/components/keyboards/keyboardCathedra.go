package keyboard

import (
	"Friends/src/assets"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboardCathedra() *telego.ReplyKeyboardMarkup {
	items := make([]telego.KeyboardButton, len(assets.Cathedras))
	for i := 0; i < len(assets.Cathedras); i++ {
		items[i] = tu.KeyboardButton(assets.Cathedras[i])
	}
	rowsCount := (len(assets.Cathedras) + 4) / 5
	items_rows := make([][]telego.KeyboardButton, rowsCount)

	for i := 0; i < rowsCount; i++ {
		start := i * 5
		end := start + 5
		if end > len(assets.Cathedras) {
			end = len(assets.Cathedras)
		}
		items_rows[i] = make([]telego.KeyboardButton, 0, end-start) // Создаем срез с нужной длиной
		for j := start; j < end; j++ {
			items_rows[i] = append(items_rows[i], tu.KeyboardButton(assets.Cathedras[j]))
		}
	}

	return tu.Keyboard(
		items_rows...,
	).WithResizeKeyboard().WithInputFieldPlaceholder("Выберите кафедру").WithOneTimeKeyboard()
}
