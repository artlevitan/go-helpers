// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/

package dates

import (
	"regexp"
	"time"
)

var (
	IsValidSQLDatePattern     = regexp.MustCompile(`^\d{4}-((0[1-9])|(1[0-2]))-((0[1-9])|([1-2][0-9])|(3[0-1]))$`)
	IsValidSQLDateTimePattern = regexp.MustCompile(`^\d{4}-((0[1-9])|(1[0-2]))-((0[1-9])|([1-2][0-9])|(3[0-1]))(\s)(([0-1][0-9])|(2[0-3])):([0-5][0-9]):([0-5][0-9])$`)
	IsValidSQLTimePattern     = regexp.MustCompile(`^(([0-1][0-9])|(2[0-3])):([0-5][0-9]):([0-5][0-9])$`)
)

// IsValidSQLDate проверяет, что строка имеет формат SQL DATE и является валидной датой
func IsValidSQLDate(d string) bool {
	if !IsValidSQLDatePattern.MatchString(d) {
		return false
	}

	// Попытка разобрать дату с использованием time.Parse
	_, err := time.Parse("2006-01-02", d)
	return err == nil
}

// IsValidSQLDateTime проверяет, что строка имеет формат SQL DATETIME и является валидной датой и временем
func IsValidSQLDateTime(d string) bool {
	if !IsValidSQLDateTimePattern.MatchString(d) {
		return false
	}

	// Попытка разобрать дату-время с использованием time.Parse
	_, err := time.Parse("2006-01-02 15:04:05", d)
	return err == nil
}

// IsValidSQLTime проверяет, что строка имеет формат SQL TIME и является валидным временем
func IsValidSQLTime(d string) bool {
	if !IsValidSQLTimePattern.MatchString(d) {
		return false
	}

	// Попытка разобрать время с использованием time.Parse
	_, err := time.Parse("15:04:05", d)
	return err == nil
}

// HasDelayPassed проверяет, превысило ли текущее время интервал времени, заданный относительно указанного начального времени startTime на добавленную задержку delay.
// Возвращает true, если текущее время позже, чем startTime + delay.
func HasDelayPassed(startTime time.Time, delay time.Duration) bool {
	return time.Now().After(startTime.Add(delay))
}

// IsTimeInRange проверяет, находится ли время check в диапазоне от start до end.
// Возвращает true, если время check находится после start и перед end.
func IsTimeInRange(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
