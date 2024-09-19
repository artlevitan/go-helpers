// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.8

package helpers

import (
	"testing"
	"time"
)

func TestHasDelayPassed(t *testing.T) {
	// Текущая метка времени
	startTime := time.Now()

	// Задержка 1 секунда
	delayDuration := 1 * time.Second

	// Ждем чуть больше 1 секунды
	time.Sleep(1100 * time.Millisecond)

	// Проверяем, что задержка прошла
	if !HasDelayPassed(startTime, delayDuration) {
		t.Errorf("Ожидалось, что HasDelayPassed вернет true, так как задержка прошла")
	}

	// Проверяем, что задержка не прошла для времени в будущем
	startTime = time.Now().Add(2 * time.Second)
	if HasDelayPassed(startTime, delayDuration) {
		t.Errorf("Ожидалось, что HasDelayPassed вернет false, так как задержка еще не прошла")
	}
}

func TestIsTimeInRange(t *testing.T) {
	// Устанавливаем диапазон времени
	startTime := time.Date(2023, 9, 15, 10, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, 9, 15, 12, 0, 0, 0, time.UTC)

	// Проверяем время внутри диапазона
	timeToCheck := time.Date(2023, 9, 15, 11, 0, 0, 0, time.UTC)
	if !IsTimeInRange(startTime, endTime, timeToCheck) {
		t.Errorf("Ожидалось, что IsTimeInRange вернет true для времени внутри диапазона")
	}

	// Проверяем время до начала диапазона
	timeToCheck = time.Date(2023, 9, 15, 9, 0, 0, 0, time.UTC)
	if IsTimeInRange(startTime, endTime, timeToCheck) {
		t.Errorf("Ожидалось, что IsTimeInRange вернет false для времени до начала диапазона")
	}

	// Проверяем время после конца диапазона
	timeToCheck = time.Date(2023, 9, 15, 13, 0, 0, 0, time.UTC)
	if IsTimeInRange(startTime, endTime, timeToCheck) {
		t.Errorf("Ожидалось, что IsTimeInRange вернет false для времени после конца диапазона")
	}

	// Проверяем, что время на границе диапазона не входит
	timeToCheck = startTime
	if IsTimeInRange(startTime, endTime, timeToCheck) {
		t.Errorf("Ожидалось, что IsTimeInRange вернет false для времени на границе начала диапазона")
	}

	timeToCheck = endTime
	if IsTimeInRange(startTime, endTime, timeToCheck) {
		t.Errorf("Ожидалось, что IsTimeInRange вернет false для времени на границе конца диапазона")
	}
}
