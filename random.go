// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.1

package helpers

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

// StringCharSet - тип для выбора набора символов
type StringCharSet int

const (
	Digits                   StringCharSet = iota // Только цифры
	Lowercase                                     // Только строчные буквы
	Uppercase                                     // Только заглавные буквы
	Letters                                       // Только буквы (заглавные и строчные)
	LettersAndDigits                              // Буквы и цифры
	LettersAndSpecials                            // Буквы и специальные символы
	LettersDigitsAndSpecials                      // Буквы, цифры и специальные символы
)

// RandomMD5 генерирует случайный MD5-хеш.
func RandomMD5() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return MD5(timestamp)
}

// RandomSHA1 генерирует случайный SHA1-хеш.
func RandomSHA1() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return SHA1(timestamp)
}

// RandomSHA256 генерирует случайный SHA256-хеш.
func RandomSHA256() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return SHA256(timestamp)
}

// RandomSHA512 генерирует случайный SHA512-хеш.
func RandomSHA512() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return SHA512(timestamp)
}

// RandomInt возвращает случайное целое число в диапазоне от min до max включительно.
func RandomInt(min, max int) int {
	// Проверяем и меняем местами min и max, если нужно
	if min > max {
		min, max = max, min
	}

	// Инициализируем новый источник случайных чисел
	source := rand.NewSource(uint64(time.Now().UnixNano()))
	r := rand.New(source)

	// Генерируем случайное число в заданном диапазоне
	return r.Intn(max-min+1) + min
}

// RandomUUID возвращает случайный UUID.
func RandomUUID() string {
	return uuid.NewString()
}

// RandomString возвращает случайную строку заданной длины из выбранного набора символов.
func RandomString(length int, charSetType StringCharSet) string {
	// Проверяем, что длина больше 0
	if length <= 0 {
		return ""
	}

	// Встроенные наборы символов
	digits := []rune("0123456789")
	lowercase := []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letters := append(uppercase, lowercase...)
	specials := []rune(":;~=+%^*()[]{}/!@#$?")

	var availableChars []rune

	// Определяем набор символов в зависимости от переданного типа
	switch charSetType {
	case Digits:
		availableChars = digits
	case Lowercase:
		availableChars = lowercase
	case Uppercase:
		availableChars = uppercase
	case Letters:
		availableChars = letters
	case LettersAndDigits:
		availableChars = append(letters, digits...)
	case LettersAndSpecials:
		availableChars = append(letters, specials...)
	case LettersDigitsAndSpecials:
		availableChars = append(letters, append(digits, specials...)...)
	}

	// Инициализируем новый источник случайных чисел
	source := rand.NewSource(uint64(time.Now().UnixNano()))
	r := rand.New(source)

	// Создаем слайс для хранения случайных символов
	buf := make([]rune, length)
	for i := range buf {
		buf[i] = availableChars[r.Intn(len(availableChars))]
	}
	return string(buf)
}
