// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.4

package helpers

import (
	"fmt"
	"math"
	"strconv"
)

// IsFloatEqual проверяет, что два числа float64 равны с учетом относительной точности.
func IsFloatEqual(num1, num2 float64) bool {
	const epsilon = 1e-9

	// Проверка на точное равенство (включая оба числа равные нулю)
	if num1 == num2 {
		return true
	}

	// Проверка на бесконечности
	if math.IsInf(num1, 0) && math.IsInf(num2, 0) {
		return math.Signbit(num1) == math.Signbit(num2)
	}

	// Проверка на NaN
	if math.IsNaN(num1) || math.IsNaN(num2) {
		return false
	}

	diff := math.Abs(num1 - num2)

	// Если хотя бы одно из чисел близко к нулю, используем абсолютную разницу
	if math.Abs(num1) < epsilon || math.Abs(num2) < epsilon {
		return diff < epsilon
	}

	// В остальных случаях используем относительную разницу
	return diff/math.Max(math.Abs(num1), math.Abs(num2)) < epsilon
}

// RoundFloat округляет число до указанного количества знаков после запятой.
func RoundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// StringToFloat преобразует строку в float64. Если строка не может быть распознана как число, возвращает 0.
func StringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil || math.IsNaN(f) {
		return 0
	}
	return f
}

// FloatToString преобразует float64 в строку с 8 знаками после запятой.
func FloatToString(f float64) string {
	return fmt.Sprintf("%.8f", f)
}

// MinOrDefault возвращает минимальное из двух чисел, или limit, если number меньше или равен нулю.
func MinOrDefault(number int, limit int) int {
	if number <= 0 {
		return limit
	}
	return min(number, limit)
}
