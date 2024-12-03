package input

import (
	"fmt"
	"regexp"
	"strings"
)

func CheckPeriod(period string) (string, error) {
	// Убираем пробелы в начале и в конце строки
	period = strings.TrimSpace(period)

	// Регулярное выражение для проверки формата "часы:минуты"
	re := regexp.MustCompile(`^([0-1]?[0-9]|2[0-3]):([0-5]?[0-9])$`)
	matches := re.FindStringSubmatch(period)

	if matches == nil {
		return "", fmt.Errorf("некорректный формат: должен быть 'час:минуты'")
	}

	// Возвращаем отформатированную строку
	return fmt.Sprintf("%02s:%02s", matches[1], matches[2]), nil
}