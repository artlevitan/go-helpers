// Copyright 2023-2025, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.2.0

package helpers

import (
	"time"
)

// HasDelayPassed проверяет, прошло ли заданное время задержки от указанного начального времени.
// Возвращает true, если текущее время позже, чем startTime + delay.
func HasDelayPassed(startTime time.Time, delayDuration time.Duration) bool {
	return time.Since(startTime) > delayDuration
}

// IsTimeInRange проверяет, находится ли заданное время в диапазоне от startTime до endTime.
// Возвращает true, если timeToCheck находится после startTime и перед endTime.
func IsTimeInRange(startTime, endTime, timeToCheck time.Time) bool {
	return timeToCheck.After(startTime) && timeToCheck.Before(endTime)
}
