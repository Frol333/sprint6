package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final-tpl/pkg/morse"
)

// DetectAndConvert определяет, является ли входная строка текстом или кодом Морзе,
// и конвертирует её в соответствующий формат.
func DetectAndConvert(input string) (string, error) {
	// Проверяем, содержит ли строка только символы Морзе (точки, тире, пробелы).
	isMorse := true
	for _, r := range input {
		if !strings.ContainsRune(".- ", r) {
			isMorse = false
			break
		}
	}

	if isMorse {
		// Если это код Морзе, конвертируем в текст.
		return morse.ToText(input), nil
	} else {
		// Если это обычный текст, конвертируем в код Морзе.
		return morse.ToMorse(input), nil
	}
}
