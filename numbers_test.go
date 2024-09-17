// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.1

package helpers

import (
	"testing"
)

func TestRoundFloat(t *testing.T) {
	type args struct {
		val       float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{val: 1.2345, precision: 2}, 1.23},
		{"2", args{val: 1.2345, precision: 3}, 1.235},
		{"3", args{val: 1.2345, precision: 0}, 1},
		{"4", args{val: 1.9999, precision: 3}, 2},
		{"5", args{val: -1.2345, precision: 2}, -1.23},
		{"6", args{val: 0.5555, precision: 2}, 0.56},
		{"7", args{val: 0.5544, precision: 2}, 0.55},
		{"8", args{val: 123456.789, precision: 1}, 123456.8},
		{"9", args{val: -0.0004, precision: 3}, 0},
		{"10", args{val: 1.1111, precision: 4}, 1.1111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundFloat(tt.args.val, tt.args.precision); got != tt.want {
				t.Errorf("RoundFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestStringToFloat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{"1.23"}, 1.23},
		{"2", args{"-1.23"}, -1.23},
		{"3", args{"0.00000001"}, 0.00000001},
		{"4", args{"123456789.123456789"}, 123456789.123456789},
		{"5", args{"0"}, 0},
		{"6", args{"NaN"}, 0},
		{"7", args{""}, 0},
		{"8", args{"abc"}, 0},
		{"9", args{"-9999.9999"}, -9999.9999},
		{"10", args{"1e10"}, 10000000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToFloat(tt.args.s); got != tt.want {
				t.Errorf("StringToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatToString(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{1.23}, "1.23000000"},
		{"2", args{-1.23}, "-1.23000000"},
		{"3", args{0.00000001}, "0.00000001"},
		{"4", args{123456789.123456789}, "123456789.12345679"},
		{"5", args{0}, "0.00000000"},
		{"6", args{-9999.9999}, "-9999.99990000"},
		{"7", args{1e10}, "10000000000.00000000"},
		{"8", args{-1e-10}, "-0.00000000"},
		{"9", args{1.111111111}, "1.11111111"},
		{"10", args{999999.99999999}, "999999.99999999"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatToString(tt.args.f); got != tt.want {
				t.Errorf("FloatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
