// Copyright 2023-2025, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.2.3

package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"reflect"
	"slices"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	siUnit  = 1000
	siUnits = "kMGTPEZY"
)

var logUnit = math.Log(siUnit)

// ByteCountSI преобразует размер файла в байтах в строку
// с использованием SI-единиц (например, kB, MB, GB).
func ByteCountSI(byteSize int64) string {
	if byteSize < siUnit {
		return fmt.Sprintf("%d B", byteSize)
	}
	exp := int(math.Log(float64(byteSize)) / logUnit)
	pre := siUnits[exp-1]
	val := float64(byteSize) / math.Pow(siUnit, float64(exp))
	return fmt.Sprintf("%.2f %cB", val, pre)
}

// ItemExists проверяет, существует ли элемент в срезе.
func ItemExists[T comparable](slice []T, item T) bool {
	return slices.Contains(slice, item)
}

// FileExists проверяет, существует ли файл и не является ли он директорией.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// Unique возвращает новый срез, содержащий только уникальные элементы.
func Unique[T comparable](inputSlice []T) []T {
	seen := make(map[T]struct{}, len(inputSlice))
	unique := make([]T, 0, len(inputSlice))
	for _, item := range inputSlice {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			unique = append(unique, item)
		}
	}
	return unique
}

// EncodeBase64 кодирует строку в Base64.
func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// DecodeBase64 декодирует строку из Base64.
// В случае ошибки возвращает пустую строку.
func DecodeBase64(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
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
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePasswords сравнивает хэшированный пароль с открытым паролем.
func ComparePasswords(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}

// CreateCacheKey создает текстовый ключ для кеширования, обрабатывая различные типы данных внутри одной функции.
func CreateCacheKey(input any) string {
	if input == nil {
		return "nil"
	}

	var sb strings.Builder

	switch v := input.(type) {
	// Обработка строк
	case string:
		sb.WriteString(v)

	// Обработка срезов байтов ([]byte или []uint8)
	case []byte:
		sb.WriteString(string(v))

	// Обработка целочисленных типов
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		_, _ = fmt.Fprintf(&sb, "%d", v)

	// Обработка срезов целочисленных типов
	case []int, []int8, []int16, []int32, []int64,
		[]uint, []uint16, []uint32, []uint64:
		rv := reflect.ValueOf(v)
		for i := range rv.Len() {
			elem := rv.Index(i).Interface()
			_, _ = fmt.Fprintf(&sb, "%d", elem)
		}

	// Обработка других типов с использованием JSON
	default:
		bytes, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprintf("error:%v", err)
		}
		sb.WriteString(string(bytes))
	}

	return sb.String()
}
