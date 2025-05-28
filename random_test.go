// Copyright 2023-2025, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.2.2

package helpers

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/google/uuid"
)

func TestRandomMD5(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test 1"}, {"test 2"}, {"test 3"}, {"test 4"}, {"test 5"},
		{"test 6"}, {"test 7"}, {"test 8"}, {"test 9"}, {"test 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomMD5()
			if len(got) != 32 {
				t.Errorf("RandomMD5() = %v, length = %d, want length 32", got, len(got))
			}
		})
	}
}

func TestRandomSHA1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test 1"}, {"test 2"}, {"test 3"}, {"test 4"}, {"test 5"},
		{"test 6"}, {"test 7"}, {"test 8"}, {"test 9"}, {"test 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomSHA1()
			if len(got) != 40 {
				t.Errorf("RandomSHA1() = %v, length = %d, want length 40", got, len(got))
			}
		})
	}
}

func TestRandomSHA256(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test 1"}, {"test 2"}, {"test 3"}, {"test 4"}, {"test 5"},
		{"test 6"}, {"test 7"}, {"test 8"}, {"test 9"}, {"test 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomSHA256()
			if len(got) != 64 {
				t.Errorf("RandomSHA256() = %v, length = %d, want length 64", got, len(got))
			}
		})
	}
}

func TestRandomSHA512(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test 1"}, {"test 2"}, {"test 3"}, {"test 4"}, {"test 5"},
		{"test 6"}, {"test 7"}, {"test 8"}, {"test 9"}, {"test 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomSHA512()
			if len(got) != 128 {
				t.Errorf("RandomSHA512() = %v, length = %d, want length 128", got, len(got))
			}
		})
	}
}

// Тест на корректный диапазон
func TestRandomIntRange(t *testing.T) {
	minVal := 10
	maxVal := 20

	for i := 0; i < 1000; i++ {
		result := RandomInt(minVal, maxVal)
		if result < minVal || result > maxVal {
			t.Errorf("Число %d вне диапазона [%d, %d]", result, minVal, maxVal)
		}
	}
}

// Тест на работу с отрицательными значениями
func TestRandomIntNegativeRange(t *testing.T) {
	minVal := -20
	maxVal := -10

	for i := 0; i < 1000; i++ {
		result := RandomInt(minVal, maxVal)
		if result < minVal || result > maxVal {
			t.Errorf("Число %d вне диапазона [%d, %d]", result, minVal, maxVal)
		}
	}
}

// Тест на корректную работу, если minVal больше maxVal
func TestRandomIntMinGreaterThanMax(t *testing.T) {
	minVal := 20
	maxVal := 10

	for i := 0; i < 1000; i++ {
		result := RandomInt(minVal, maxVal)
		if result < maxVal || result > minVal {
			t.Errorf("Число %d вне диапазона [%d, %d] при minVal > maxVal", result, maxVal, minVal)
		}
	}
}

// Тест на одинаковые значения minVal и maxVal
func TestRandomIntMinEqualsMax(t *testing.T) {
	minVal := 10
	maxVal := 10

	for i := 0; i < 100; i++ {
		result := RandomInt(minVal, maxVal)
		if result != minVal {
			t.Errorf("При minVal = maxVal ожидается %d, получено %d", minVal, result)
		}
	}
}

// Тест на корректную генерацию чисел за короткое время
func TestRandomIntTimeVariance(t *testing.T) {
	minVal := 1
	maxVal := 5

	// Проверим, что сгенерированные числа разные при многократном вызове функции
	results := make(map[int]bool)
	for i := 0; i < 100; i++ {
		result := RandomInt(minVal, maxVal)
		results[result] = true
	}

	if len(results) == 1 {
		t.Errorf("Случайные числа не меняются при многократных вызовах")
	}
}

// Тест на проверку корректного формата UUID
func TestRandomUUIDFormat(t *testing.T) {
	u := RandomUUID()
	_, err := uuid.Parse(u)
	if err != nil {
		t.Errorf("Неверный формат UUID: %s", u)
	}
}

// Тест на уникальность UUID при многократном вызове
func TestRandomUUIDUniqueness(t *testing.T) {
	iterations := 1000
	uuidSet := make(map[string]struct{})

	for i := 0; i < iterations; i++ {
		u := RandomUUID()
		if _, exists := uuidSet[u]; exists {
			t.Errorf("Найден дубликат UUID: %s", u)
		}
		uuidSet[u] = struct{}{}
	}
}

// Тест на корректную длину возвращаемой строки
func TestRandomStringLength(t *testing.T) {
	length := 10
	result := RandomString(length, Letters)
	if len(result) != length {
		t.Errorf("Ожидалось, что длина строки будет %d, но получено %d", length, len(result))
	}
}

// Тест на возврат пустой строки при нулевой длине
func TestRandomStringZeroLength(t *testing.T) {
	result := RandomString(0, Letters)
	if result != "" {
		t.Errorf("Ожидалась пустая строка при длине 0, но получено: %s", result)
	}
}

// Тест на возврат пустой строки при отрицательной длине
func TestRandomStringNegativeLength(t *testing.T) {
	result := RandomString(-5, Letters)
	if result != "" {
		t.Errorf("Ожидалась пустая строка при отрицательной длине, но получено: %s", result)
	}
}

// Тест на генерацию строки только из цифр
func TestRandomStringDigits(t *testing.T) {
	result := RandomString(10, Digits)
	for _, char := range result {
		if !unicode.IsDigit(char) {
			t.Errorf("Ожидалась строка, содержащая только цифры, но найден символ: %c", char)
		}
	}
}

// Тест на генерацию строки только из строчных букв
func TestRandomStringLowercase(t *testing.T) {
	result := RandomString(10, Lowercase)
	for _, char := range result {
		if !unicode.IsLower(char) {
			t.Errorf("Ожидалась строка, содержащая только строчные буквы, но найден символ: %c", char)
		}
	}
}

// Тест на генерацию строки только из заглавных букв
func TestRandomStringUppercase(t *testing.T) {
	result := RandomString(10, Uppercase)
	for _, char := range result {
		if !unicode.IsUpper(char) {
			t.Errorf("Ожидалась строка, содержащая только заглавные буквы, но найден символ: %c", char)
		}
	}
}

// Тест на генерацию строки из букв и цифр
func TestRandomStringLettersAndDigits(t *testing.T) {
	result := RandomString(10, LettersAndDigits)
	for _, char := range result {
		if !(unicode.IsLetter(char) || unicode.IsDigit(char)) {
			t.Errorf("Ожидалась строка, содержащая только буквы и цифры, но найден символ: %c", char)
		}
	}
}

// Тест на генерацию строки с буквами, цифрами и специальными символами
func TestRandomStringLettersDigitsAndSpecials(t *testing.T) {
	specials := ":;~=+%^*()[]{}/!@#$?"
	result := RandomString(50, LettersDigitsAndSpecials)
	valid := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + specials
	for _, char := range result {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && !containsRune(valid, char) {
			t.Errorf("Ожидалась строка, содержащая буквы, цифры или специальные символы, но найден символ: %c", char)
		}
	}
}

// Тест на генерацию строки с буквами и специальными символами
func TestRandomStringLettersAndSpecials(t *testing.T) {
	specials := ":;~=+%^*()[]{}/!@#$?"
	result := RandomString(50, LettersAndSpecials)
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	valid := letters + specials

	for _, char := range result {
		if !unicode.IsLetter(char) && !containsRune(valid, char) {
			t.Errorf("Ожидалась строка, содержащая буквы или специальные символы, но найден символ: %c", char)
		}
	}
}

// Вспомогательная функция для проверки наличия символа в строке
func containsRune(str string, r rune) bool {
	for _, c := range str {
		if c == r {
			return true
		}
	}
	return false
}

func TestRandomCodeEdgeCases(t *testing.T) {
	cases := []struct {
		length    int
		wantEmpty bool
	}{
		{length: -5, wantEmpty: true},
		{length: 0, wantEmpty: true},
		{length: 1, wantEmpty: false},
		{length: 5, wantEmpty: false},
		{length: 10, wantEmpty: false},
	}

	for _, tc := range cases {
		t.Run(
			fmt.Sprintf("length=%d", tc.length),
			func(t *testing.T) {
				got := RandomCode(tc.length)
				if tc.wantEmpty {
					if got != "" {
						t.Errorf("RandomCode(%d) = %q; want empty string", tc.length, got)
					}
					return
				}
				// Длина должна совпадать
				if len(got) != tc.length {
					t.Errorf("RandomCode(%d) length = %d; want %d", tc.length, len(got), tc.length)
				}
				// Все символы — цифры
				for i, r := range got {
					if !unicode.IsDigit(r) {
						t.Errorf("RandomCode(%d)[%d] = %q; want digit", tc.length, i, r)
					}
				}
			},
		)
	}
}

func TestRandomCodeVariability(t *testing.T) {
	const length = 5
	seen := make(map[string]struct{})
	const iterations = 1000

	for i := 0; i < iterations; i++ {
		code := RandomCode(length)
		seen[code] = struct{}{}
	}
	// Ожидаем, что хотя бы несколько кодов различаются
	if len(seen) < 10 {
		t.Errorf("RandomCode(%d) variability too low: only %d unique codes out of %d calls", length, len(seen), iterations)
	}
}
