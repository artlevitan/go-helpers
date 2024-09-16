// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/

package strings

import (
	"regexp"
	"strings"
)

var (
	clearStringPattern   = regexp.MustCompile(`\s+`)
	filterLettersPattern = regexp.MustCompile(`[^a-zA-Zа-яА-Я]+`)
	filterDigitsPattern  = regexp.MustCompile(`\D+`)
)

// CutString обрезает строку до заданной длины.
func CutString(text string, length int) string {
	runes := []rune(text)
	if len(runes) > length {
		return string(runes[:length])
	}
	return text
}

// ClearString убирает пробелы с краев строки и заменяет дублирующие пробелы одним пробелом.
func ClearString(text string) string {
	text = clearStringPattern.ReplaceAllString(text, " ")
	return strings.TrimSpace(text)
}

// FilterLetters очищает строку от всех символов, кроме букв.
// Возвращает очищенную строку.
func FilterLetters(text string) string {
	text = strings.TrimSpace(text)
	return filterLettersPattern.ReplaceAllString(text, "")
}

// FilterDigits очищает строку от всех символов, кроме цифр.
// Возвращает очищенную строку.
func FilterDigits(text string) string {
	text = strings.TrimSpace(text)
	return filterDigitsPattern.ReplaceAllString(text, "")
}
