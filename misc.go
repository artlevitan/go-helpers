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
