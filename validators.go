// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.1

package helpers

import (
	"encoding/json"
	"net"
	"net/url"
	"regexp"
	"time"
)

var (
	isValidSQLDatePattern     = regexp.MustCompile(`^\d{4}-((0[1-9])|(1[0-2]))-((0[1-9])|([1-2][0-9])|(3[0-1]))$`)
	isValidSQLDateTimePattern = regexp.MustCompile(`^\d{4}-((0[1-9])|(1[0-2]))-((0[1-9])|([1-2][0-9])|(3[0-1]))(\s)(([0-1][0-9])|(2[0-3])):([0-5][0-9]):([0-5][0-9])$`)
	isValidSQLTimePattern     = regexp.MustCompile(`^(([0-1][0-9])|(2[0-3])):([0-5][0-9]):([0-5][0-9])$`)
	isValidHexColorPattern    = regexp.MustCompile(`^#?([A-Fa-f\d]{3}|[A-Fa-f\d]{6})$`)
)

// IsValidSQLDate проверяет, что строка имеет формат SQL DATE и является валидной датой.
func IsValidSQLDate(d string) bool {
	if !isValidSQLDatePattern.MatchString(d) {
		return false
	}

	// Попытка разобрать дату с использованием time.Parse
	_, err := time.Parse("2006-01-02", d)
	return err == nil
}

// IsValidSQLDateTime проверяет, что строка имеет формат SQL DATETIME и является валидной датой и временем.
func IsValidSQLDateTime(d string) bool {
	if !isValidSQLDateTimePattern.MatchString(d) {
		return false
	}

	// Попытка разобрать дату-время с использованием time.Parse
	_, err := time.Parse("2006-01-02 15:04:05", d)
	return err == nil
}

// IsValidSQLTime проверяет, что строка имеет формат SQL TIME и является валидным временем.
func IsValidSQLTime(d string) bool {
	if !isValidSQLTimePattern.MatchString(d) {
		return false
	}

	// Попытка разобрать время с использованием time.Parse
	_, err := time.Parse("15:04:05", d)
	return err == nil
}

// IsValidHexColor проверяет, является ли строка валидным HEX-кодом цвета.
func IsValidHexColor(color string) bool {
	return isValidHexColorPattern.MatchString(color)
}

// IsURL проверяет, является ли переданная строка валидной URL-ссылкой.
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && len(u.Scheme) > 0 && len(u.Host) > 0
}

// IsIPv4 проверяет, является ли предоставленная строка действительным IPv4-адресом.
func IsIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
}

// IsIPv6 проверяет, является ли предоставленная строка действительным IPv6-адресом.
func IsIPv6(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To16() != nil && parsedIP.To4() == nil
}

// IsValidJSON проверяет, является ли строка валидным JSON.
func IsValidJSON(s string) (bool, error) {
	var js interface{}
	err := json.Unmarshal([]byte(s), &js)
	if err != nil {
		return false, err
	}
	return true, nil
}
