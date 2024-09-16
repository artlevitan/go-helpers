// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/

package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
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

// CheckStringLength проверяет, что длина строки находится в заданном диапазоне.
// Возвращает true, если длина строки в диапазоне [minLength, maxLength], иначе false.
func CheckStringLength(str string, minLength int, maxLength int) bool {
	length := len([]rune(str))
	return length >= minLength && length <= maxLength
}

// CreateCacheKey создает текстовый ключ для кеширования, обрабатывая различные типы данных внутри одной функции.
func CreateCacheKey(in interface{}) string {
	if in == nil {
		return "nil"
	}

	var sb strings.Builder

	switch v := in.(type) {
	// Обработка строк
	case string:
		sb.WriteString(v)

	// Обработка срезов байтов ([]byte или []uint8)
	case []byte:
		sb.WriteString(string(v))

	// Обработка знаковых целочисленных типов
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		fmt.Fprintf(&sb, "%d", v)

	// Обработка срезов целочисленных типов
	case []int, []int8, []int16, []int32, []int64,
		[]uint, []uint16, []uint32, []uint64:
		// Используем рефлексию для обработки срезов различных типов
		rv := reflect.ValueOf(v)
		for i := 0; i < rv.Len(); i++ {
			elem := rv.Index(i).Interface()
			fmt.Fprintf(&sb, "%d", elem)
		}

	// Обработка других типов с использованием JSON-сериализации
	default:
		bytes, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprintf("error:%v", err)
		}
		sb.WriteString(string(bytes))
	}

	return sb.String()
}
