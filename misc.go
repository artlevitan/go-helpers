// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.3

package helpers

import (
	"encoding/base64"
	"fmt"
	"math"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// ByteCountSI преобразует размер файла в байтах в строку
// с использованием SI-единиц (например, kB, MB, GB).
func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	exp := int(math.Log(float64(b)) / math.Log(unit))
	pre := "kMGTPEZY"[exp-1]
	value := float64(b) / math.Pow(unit, float64(exp))
	return fmt.Sprintf("%.2f %cB", value, pre)
}

// ItemExists проверяет, существует ли элемент в срезе.
func ItemExists[T comparable](slice []T, item T) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

// FileExists проверяет, существует ли файл и не является ли он директорией.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Unique возвращает новый срез, содержащий только уникальные элементы.
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

// ActiveEnum возвращает значение ENUM["0", "1"] в зависимости от входного флага.
func ActiveEnum(flag string) string {
	if flag == "1" {
		return "1"
	}
	return "0"
}

// HashPassword хэширует пароль.
func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePasswords сравнивает хэшированный пароль с plain текстом.
func ComparePasswords(hashedPwd, plainPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)) == nil
}
