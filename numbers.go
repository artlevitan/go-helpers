// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.1

package helpers

import (
	"fmt"
	"math"
	"strconv"
)

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
