// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.3

package helpers

import (
	"time"
)

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
