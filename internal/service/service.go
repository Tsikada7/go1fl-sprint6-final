package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Функция определяет тип ввода и переводит его либо в текст либо морзе
func Converter(input string) (string, error) {

	if strings.TrimSpace(input) == "" {
		return "", errors.New("empty input")
	}
	morseOnly := strings.Trim(input, ".- \n\r")
	if morseOnly == "" {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}
