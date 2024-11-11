package utils

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func GetItemsRow(data []string) [][]telego.KeyboardButton {
	items := make([]telego.KeyboardButton, len(data))
	for i := 0; i < len(data); i++ {
		items[i] = tu.KeyboardButton(data[i])
	}
	rowsCount := (len(data) + 4) / 5
	items_rows := make([][]telego.KeyboardButton, rowsCount)

	for i := 0; i < rowsCount; i++ {
		start := i * 5
		end := start + 5
		if end > len(data) {
			end = len(data)
		}
		items_rows[i] = make([]telego.KeyboardButton, 0, end-start)
		for j := start; j < end; j++ {
			items_rows[i] = append(items_rows[i], tu.KeyboardButton(data[j]))
		}
	}
	return items_rows
}
