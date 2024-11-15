// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.15

package helpers

import (
	"encoding/json"
	"net"
	"net/url"
	"regexp"
	"time"
)

var (
	isSQLDatePattern     = regexp.MustCompile(`^\d{4}-((0[1-9])|(1[0-2]))-((0[1-9])|([1-2][0-9])|(3[0-1]))$`)
	isSQLDateTimePattern = regexp.MustCompile(`^\d{4}-((0[1-9])|(1[0-2]))-((0[1-9])|([1-2][0-9])|(3[0-1]))(\s)(([0-1][0-9])|(2[0-3])):([0-5][0-9]):([0-5][0-9])$`)
	isSQLTimePattern     = regexp.MustCompile(`^(([0-1][0-9])|(2[0-3])):([0-5][0-9]):([0-5][0-9])$`)
	isHexColorPattern    = regexp.MustCompile(`^#?([A-Fa-f\d]{3}|[A-Fa-f\d]{6})$`)
)

// IsSQLDate проверяет, что строка имеет формат SQL DATE и является валидной датой.
func IsSQLDate(d string) bool {
	if !isSQLDatePattern.MatchString(d) {
		return false
	}

	// Попытка разобрать строку
	_, err := time.Parse("2006-01-02", d)
	return err == nil
}

// IsSQLDateTime проверяет, что строка имеет формат SQL DATETIME и является валидной датой и временем.
func IsSQLDateTime(d string) bool {
	if !isSQLDateTimePattern.MatchString(d) {
		return false
	}

	// Попытка разобрать строку
	_, err := time.Parse("2006-01-02 15:04:05", d)
	return err == nil
}

// IsSQLTime проверяет, что строка имеет формат SQL TIME и является валидным временем.
func IsSQLTime(d string) bool {
	if !isSQLTimePattern.MatchString(d) {
		return false
	}

	// Попытка разобрать строку
	_, err := time.Parse("15:04:05", d)
	return err == nil
}

// IsHexColor проверяет, является ли строка валидным HEX-кодом цвета.
func IsHexColor(color string) bool {
	return isHexColorPattern.MatchString(color)
}

// IsURL проверяет, является ли переданная строка валидной URL-ссылкой.
func IsURL(s string) bool {
	u, err := url.Parse(s)
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

// IsPrivateOrReservedIP проверяет, является ли указанный IP-адрес частным или зарезервированным.
func IsPrivateOrReservedIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	// Проверка частных, loopback, unspecified, multicast и link-local адресов
	if parsedIP.IsPrivate() || parsedIP.IsLoopback() || parsedIP.IsUnspecified() ||
		parsedIP.IsMulticast() || parsedIP.IsLinkLocalUnicast() {
		return true
	}

	// Проверка IPv6 диапазона документации (2001:db8::/32)
	if parsedIP.To16() != nil {
		_, db8Net, _ := net.ParseCIDR("2001:db8::/32")
		if db8Net.Contains(parsedIP) {
			return true
		}
	}

	// Проверка IPv4 диапазонов документации
	if parsedIP.To4() != nil {
		docNets := []string{
			"192.0.2.0/24",
			"198.51.100.0/24",
			"203.0.113.0/24",
		}
		for _, cidr := range docNets {
			_, net, _ := net.ParseCIDR(cidr)
			if net.Contains(parsedIP) {
				return true
			}
		}
	}

	return false
}

// IsJSON проверяет, является ли строка валидным JSON.
func IsJSON(s string) (bool, error) {
	var js interface{}
	err := json.Unmarshal([]byte(s), &js)
	if err != nil {
		return false, err
	}
	return true, nil
}
