// Copyright 2023-2025, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.2.3

package helpers

import (
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
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

// StringLength возвращает количество символов в строке.
func StringLength(s string) int {
	return utf8.RuneCountInString(s)
}

// IsStringLengthInRange проверяет, что длина строки находится в заданном диапазоне.
// Возвращает true, если длина строки в диапазоне [minLength, maxLength], иначе false.
func IsStringLengthInRange(s string, minLength int, maxLength int) bool {
	length := StringLength(s)
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

// ClearPhone форматирует номер телефона и результат обрезается до 25 символов.
func ClearPhone(text string) string {
	digits := FilterDigits(text)
	return CutString(digits, 25)
}

// ClearEmail форматирует Email, удаляя пробелы и проверяя адрес; при ошибке возвращает пустую строку.
func ClearEmail(email string) string {
	email = strings.TrimSpace(email)
	e, err := mail.ParseAddress(email)
	if err != nil {
		return ""
	}
	return e.Address
}

// Capitalize приводит строку s к формату: первая буква — в верхнем регистре, остальные — в нижнем.
func Capitalize(s string) string {
	// Если строка пустая, возвращаем ее без изменений
	if s == "" {
		return s
	}

	// Преобразуем строку в срез рун для корректной работы с символами Unicode
	runes := []rune(s)

	// Создаем strings.Builder
	var b strings.Builder
	b.Grow(len(runes))

	// Преобразуем первую руну (символ) к верхнему регистру и записываем в билдер
	b.WriteRune(unicode.ToUpper(runes[0]))

	// Обрабатываем все оставшиеся руны, преобразуя их к нижнему регистру
	for _, r := range runes[1:] {
		b.WriteRune(unicode.ToLower(r))
	}

	return b.String()
}

// CreateInitials приводит Фамилию Имя Отчество к формату "Фамилия И.О."
func CreateInitials(rSurname, rName string, rPatronymic *string) string {
	// Приводим фамилию к корректному регистру с помощью функции Capitalize
	surname := Capitalize(rSurname)

	// Создаем strings.Builder
	var b strings.Builder
	b.WriteString(surname) // Добавляем фамилию
	b.WriteString(" ")     // Добавляем пробел между фамилией и инициалами

	// Добавляем первую букву имени, приводя её к верхнему регистру
	b.WriteString(strings.ToUpper(CutString(rName, 1)))
	b.WriteString(".") // Завершаем инициалы имени точкой.

	// Если отчество передано и не является пустой строкой,
	// добавляем его первую букву, также приводя к верхнему регистру
	if rPatronymic != nil && *rPatronymic != "" {
		b.WriteString(strings.ToUpper(CutString(*rPatronymic, 1)))
		b.WriteString(".") // Завершаем инициалы отчества точкой
	}

	return b.String()
}

// NormalizeFilename принимает имя файла и возвращает его URL-совместимое представление.
func NormalizeFilename(filename string) string {
	// Заменяем пробелы на нижнее подчеркивание
	filename = strings.ReplaceAll(filename, " ", "_")

	// Экранируем строку для безопасного использования в URL
	return url.PathEscape(filename)
}
