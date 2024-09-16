// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.0

package helpers

import (
	"testing"
	"time"
)

func TestHasDelayPassed(t *testing.T) {
	// Текущая метка времени
	startTime := time.Now()

	// Задержка 1 секунда
	delay := 1 * time.Second

	// Ждем чуть больше 1 секунды
	time.Sleep(1100 * time.Millisecond)

	// Проверяем, что задержка прошла
	if !HasDelayPassed(startTime, delay) {
		t.Errorf("Ожидалось, что HasDelayPassed вернет true, так как задержка прошла")
	}

	// Проверяем, что задержка не прошла для времени в будущем
	startTime = time.Now().Add(2 * time.Second)
	if HasDelayPassed(startTime, delay) {
		t.Errorf("Ожидалось, что HasDelayPassed вернет false, так как задержка еще не прошла")
	}
}

func TestIsTimeInRange(t *testing.T) {
	// Устанавливаем диапазон времени
	start := time.Date(2023, 9, 15, 10, 0, 0, 0, time.UTC)
	end := time.Date(2023, 9, 15, 12, 0, 0, 0, time.UTC)

	// Проверяем время внутри диапазона
	checkTime := time.Date(2023, 9, 15, 11, 0, 0, 0, time.UTC)
	if !IsTimeInRange(start, end, checkTime) {
		t.Errorf("Ожидалось, что IsTimeInRange вернет true для времени внутри диапазона")
	}

	// Проверяем время до начала диапазона
	checkTime = time.Date(2023, 9, 15, 9, 0, 0, 0, time.UTC)
	if IsTimeInRange(start, end, checkTime) {
		t.Errorf("Ожидалось, что IsTimeInRange вернет false для времени до начала диапазона")
	}

	// Проверяем время после конца диапазона
	checkTime = time.Date(2023, 9, 15, 13, 0, 0, 0, time.UTC)
	if IsTimeInRange(start, end, checkTime) {
		t.Errorf("Ожидалось, что IsTimeInRange вернет false для времени после конца диапазона")
	}

	// Проверяем, что время на границе диапазона не входит
	checkTime = start
	if IsTimeInRange(start, end, checkTime) {
		t.Errorf("Ожидалось, что IsTimeInRange вернет false для времени на границе начала диапазона")
	}

	checkTime = end
	if IsTimeInRange(start, end, checkTime) {
		t.Errorf("Ожидалось, что IsTimeInRange вернет false для времени на границе конца диапазона")
	}
}
