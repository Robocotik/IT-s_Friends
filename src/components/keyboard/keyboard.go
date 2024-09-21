package keyboard

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

func CreateKeyboard() tu.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow( // Row 1
			tu.KeyboardButton("Button"),
			tu.KeyboardButton("Poll Regular").WithRequestPoll(tu.PollTypeRegular()),
		),
		tu.KeyboardRow( // Row 2
			tu.KeyboardButton("Contact").WithRequestContact(),
			tu.KeyboardButton("Location").WithRequestLocation(),
		),
	).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")
}
