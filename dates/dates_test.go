// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/

package dates

import (
	"testing"
	"time"
)

func TestIsValidSQLDate(t *testing.T) {
	// Валидные даты
	validDates := []string{
		"2023-01-01",
		"1999-12-31",
		"2024-02-29", // високосный год
		"2020-06-15",
	}

	// Невалидные даты
	invalidDates := []string{
		"2023-13-01", // неверный месяц
		"2023-00-10", // неверный месяц
		"2023-12-32", // неверный день
		"2023-02-30", // неверный день
		"2001-02-29", // не високосный год
		"2023/12/01", // неправильный формат
		"23-12-01",   // недостаточно цифр в годе
	}

	for _, date := range validDates {
		if !IsValidSQLDate(date) {
			t.Errorf("Ожидалось, что дата %s будет валидной", date)
		}
	}

	for _, date := range invalidDates {
		if IsValidSQLDate(date) {
			t.Errorf("Ожидалось, что дата %s будет невалидной", date)
		}
	}
}

func TestIsValidSQLDateTime(t *testing.T) {
	// Валидные даты-время
	validDateTimes := []string{
		"2023-01-01 12:30:45",
		"1999-12-31 23:59:59",
		"2020-02-29 00:00:00", // високосный год
		"2022-08-15 15:45:30",
	}

	// Невалидные даты-время
	invalidDateTimes := []string{
		"2023-01-01 24:00:00", // неверное время
		"2023-12-01 23:60:00", // неверная минута
		"2023-12-01 12:30:60", // неверная секунда
		"2023-13-01 12:30:45", // неверный месяц
		"2023-12-32 12:30:45", // неверный день
		"2023-12-01T12:30:45", // неправильный формат
		"2023-01-01 12:30",    // недостаточно времени
	}

	for _, dateTime := range validDateTimes {
		if !IsValidSQLDateTime(dateTime) {
			t.Errorf("Ожидалось, что дата-время %s будет валидной", dateTime)
		}
	}

	for _, dateTime := range invalidDateTimes {
		if IsValidSQLDateTime(dateTime) {
			t.Errorf("Ожидалось, что дата-время %s будет невалидной", dateTime)
		}
	}
}

func TestIsValidSQLTime(t *testing.T) {
	// Валидное время
	validTimes := []string{
		"00:00:00",
		"23:59:59",
		"12:30:45",
		"07:45:00",
	}

	// Невалидное время
	invalidTimes := []string{
		"24:00:00",   // неверный час
		"12:60:00",   // неверная минута
		"12:00:60",   // неверная секунда
		"12:00",      // недостаточно времени
		"12:00:00:1", // лишняя секция
		"120000",     // отсутствуют разделители
	}

	for _, timeStr := range validTimes {
		if !IsValidSQLTime(timeStr) {
			t.Errorf("Ожидалось, что время %s будет валидным", timeStr)
		}
	}

	for _, timeStr := range invalidTimes {
		if IsValidSQLTime(timeStr) {
			t.Errorf("Ожидалось, что время %s будет невалидным", timeStr)
		}
	}
}

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
