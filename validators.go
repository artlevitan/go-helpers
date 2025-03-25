// Copyright 2023-2025, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.2.1

package helpers

import (
	"encoding/json"
	"net/netip"
	"net/url"
	"regexp"
	"time"
)

const (
	sqlDateLayout     = time.DateOnly
	sqlDateTimeLayout = time.DateTime
	sqlTimeLayout     = time.TimeOnly
)

var (
	isSQLDatePattern     = regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`)
	isSQLDateTimePattern = regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])\s((0\d)|(1\d)|(2[0-3])):([0-5]\d):([0-5]\d)$`)
	isSQLTimePattern     = regexp.MustCompile(`^((0\d)|(1\d)|(2[0-3])):([0-5]\d):([0-5]\d)$`)
	isHexColorPattern    = regexp.MustCompile(`^#?([A-Fa-f\d]{3}|[A-Fa-f\d]{6})$`)

	// Глобальные переменные для диапазонов зарезервированных IP-адресов.
	// Диапазон для IPv6: документация (2001:db8::/32).
	docIPv6Prefix = netip.MustParsePrefix("2001:db8::/32")
	// Диапазоны для IPv4: документация.
	reservedIPv4Prefixes = []netip.Prefix{
		netip.MustParsePrefix("192.0.2.0/24"),
		netip.MustParsePrefix("198.51.100.0/24"),
		netip.MustParsePrefix("203.0.113.0/24"),
	}
)

// IsSQLDate проверяет, что строка имеет формат SQL DATE и является валидной датой.
func IsSQLDate(d string) bool {
	if !isSQLDatePattern.MatchString(d) {
		return false
	}
	// Разбор строки с использованием константного формата
	_, err := time.Parse(sqlDateLayout, d)
	return err == nil
}

// IsSQLDateTime проверяет, что строка имеет формат SQL DATETIME и является валидной датой и временем.
func IsSQLDateTime(d string) bool {
	if !isSQLDateTimePattern.MatchString(d) {
		return false
	}
	_, err := time.Parse(sqlDateTimeLayout, d)
	return err == nil
}

// IsSQLTime проверяет, что строка имеет формат SQL TIME и является валидным временем.
func IsSQLTime(d string) bool {
	if !isSQLTimePattern.MatchString(d) {
		return false
	}
	_, err := time.Parse(sqlTimeLayout, d)
	return err == nil
}

// IsHexColor проверяет, является ли строка валидным HEX-кодом цвета.
func IsHexColor(color string) bool {
	return isHexColorPattern.MatchString(color)
}

// IsURL проверяет, является ли переданная строка валидной URL-ссылкой.
func IsURL(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// IsIPv4 проверяет, является ли предоставленная строка действительным IPv4-адресом.
func IsIPv4(ip string) bool {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false
	}
	return addr.Is4()
}

// IsIPv6 проверяет, является ли предоставленная строка действительным IPv6-адресом.
func IsIPv6(ip string) bool {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false
	}
	return addr.Is6()
}

// IsPrivateOrReservedIP проверяет, является ли указанный IP-адрес частным или зарезервированным.
func IsPrivateOrReservedIP(ip string) bool {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return false
	}

	// Используем встроенные методы netip.Addr для базовой проверки
	if addr.IsPrivate() || addr.IsLoopback() || addr.IsUnspecified() ||
		addr.IsMulticast() || addr.IsLinkLocalUnicast() {
		return true
	}

	// Для IPv6: проверка диапазона документации 2001:db8::/32
	if addr.Is6() && docIPv6Prefix.Contains(addr) {
		return true
	}

	// Для IPv4: проверка диапазонов документации
	if addr.Is4() {
		for _, prefix := range reservedIPv4Prefixes {
			if prefix.Contains(addr) {
				return true
			}
		}
	}

	return false
}

// IsJSON проверяет, является ли строка валидным JSON.
func IsJSON(s string) (bool, error) {
	var js any
	err := json.Unmarshal([]byte(s), &js)
	if err != nil {
		return false, err
	}
	return true, nil
}
