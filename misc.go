// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/

package helpers

import "encoding/base64"

// ItemExists проверяет, существует ли элемент в срезе.
func ItemExists[T comparable](slice []T, item T) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

// Unique возвращает новый срез, содержащий только уникальные элементы из inputSlice
func Unique[T comparable](inputSlice []T) []T {
	seen := make(map[T]bool)
	unique := []T{}

	for _, item := range inputSlice {
		if !seen[item] {
			seen[item] = true
			unique = append(unique, item)
		}
	}

	return unique
}

// EncodeBase64 шифрует строку в Base64 формат.
func EncodeBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

// DecodeBase64 расшифровывает строку из Base64 формата.
// В случае ошибки возвращает пустую строку.
func DecodeBase64(text string) string {
	decoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}
	return string(decoded)
}
