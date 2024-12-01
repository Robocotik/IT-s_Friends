package utils

func GetChZnPhrases(del string) []string {
	isCh := del == "числитель" // переписать на словарь
	phrase_ch := "Числитель"
	phrase_zn := "Знаменатель"
	if isCh {
		phrase_ch += " (Сегодня)"
	} else {
		phrase_zn += " (Сегодня)"
	}
	return []string{phrase_ch, phrase_zn}
}
