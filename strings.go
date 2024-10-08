// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.10

package helpers

import (
	"regexp"
	"strings"
)

var (
	clearStringPattern         = regexp.MustCompile(`\s+`)
	clearTextareaPattern       = regexp.MustCompile(`(?:\r?\n){2,}|\r?\n(\s)*`)
	filterLettersPattern       = regexp.MustCompile(`(?i)[^a-zа-яё]+`)
	filterDigitsPattern        = regexp.MustCompile(`\D+`)
	clearHtmlPageStylesPattern = regexp.MustCompile(`(?i)(?:<style>((?:.*?\r?\n?)*)</style>)+|(?:<script>((?:.*?\r?\n?)*)</script>)+|\s+(class="(.*?)")|\s+(style="(.*?)")`)
	stripTagsPattern           = regexp.MustCompile(`<[^>]+>|<!--.*?-->`)
	emojiPattern               = regexp.MustCompile(`[\x{1F600}-\x{1F64F}]|` + // Emoticons
		`[\x{1F300}-\x{1F5FF}]|` + // Misc Symbols and Pictographs
		`[\x{1F680}-\x{1F6FF}]|` + // Transport and Map Symbols
		`[\x{2600}-\x{26FF}]|` + // Misc symbols
		`[\x{2700}-\x{27BF}]|` + // Dingbats
		`[\x{FE00}-\x{FE0F}]|` + // Variation Selectors
		`[\x{1F900}-\x{1F9FF}]|` + // Supplemental Symbols and Pictographs
		`[\x{1F1E6}-\x{1F1FF}]|` + // Flags (regional indicators)
		`[\x{1FA70}-\x{1FAFF}]|` + // Symbols and Pictographs Extended-A
		`[\x{200D}]`)
)

// CutString обрезает строку до заданной длины.
func CutString(s string, length int) string {
	runes := []rune(s)
	if len(runes) > length {
		return string(runes[:length])
	}
	return s
}

// ClearString убирает пробелы с краев строки и заменяет дублирующие пробелы одним пробелом.
func ClearString(s string) string {
	s = clearStringPattern.ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}

// ClearTextarea форматирует и очищает многострочный текст.
func ClearTextarea(s string) string {
	s = strings.TrimSpace(s)
	return clearTextareaPattern.ReplaceAllString(s, "\n")
}

// FilterLetters очищает строку от всех символов, кроме букв.
// Возвращает очищенную строку.
func FilterLetters(s string) string {
	s = strings.TrimSpace(s)
	return filterLettersPattern.ReplaceAllString(s, "")
}

// FilterDigits очищает строку от всех символов, кроме цифр.
// Возвращает очищенную строку.
func FilterDigits(s string) string {
	s = strings.TrimSpace(s)
	return filterDigitsPattern.ReplaceAllString(s, "")
}

// CheckStringLength проверяет, что длина строки находится в заданном диапазоне.
// Возвращает true, если длина строки в диапазоне [minLength, maxLength], иначе false.
func CheckStringLength(s string, minLength int, maxLength int) bool {
	length := len([]rune(s))
	return length >= minLength && length <= maxLength
}

// SanitizeHTML очищает строку от HTML стилей, тегов и скриптов.
func SanitizeHTML(s string) string {
	// Удаление HTML стилей, скриптов
	s = clearHtmlPageStylesPattern.ReplaceAllString(s, "")

	// Удаление HTML тегов
	s = stripTagsPattern.ReplaceAllString(s, "")

	// Очистка строки от лишних пробелов и переносов строк
	return ClearString(s)
}

// SanitizeHTMLWithTextarea очищает строку от HTML стилей, тегов и скриптов сохраня переносы.
func SanitizeHTMLWithTextarea(s string) string {
	// Удаление HTML стилей, скриптов
	s = clearHtmlPageStylesPattern.ReplaceAllString(s, "")

	// Удаление HTML тегов
	s = stripTagsPattern.ReplaceAllString(s, "")

	// Очищает многострочный текст
	return ClearTextarea(s)
}

// RemoveEmojis удаляет Emoji из строки.
func RemoveEmojis(s string) string {
	s = emojiPattern.ReplaceAllString(s, "")

	// Очистка строки от лишних пробелов и переносов строк
	return ClearString(s)
}
