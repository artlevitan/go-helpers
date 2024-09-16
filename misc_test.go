// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.0

package helpers

import (
	"testing"
)

func TestByteCountSI(t *testing.T) {
	type args struct {
		b int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{0}, "0 B"},
		{"2", args{999}, "999 B"},
		{"3", args{1023}, "1.02 kB"},
		{"4", args{1024}, "1.02 kB"},
		{"5", args{8192}, "8.19 kB"},
		{"6", args{1000000}, "1.00 MB"},
		{"7", args{1500000}, "1.50 MB"},
		{"8", args{1000000000}, "1.00 GB"},
		{"9", args{1500000000}, "1.50 GB"},
		{"10", args{1000000000000}, "1.00 TB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteCountSI(tt.args.b); got != tt.want {
				t.Errorf("ByteCountSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemExists(t *testing.T) {
	tests := []struct {
		name   string
		slice  []string
		item   string
		expect bool
	}{
		{"1", []string{"no", "day", "week", "month", "year"}, "month", true},
		{"2", []string{"no", "day", "week", "month", "year"}, "days", false},
		{"3", []string{"no", "day", "week", "month", "year"}, " year", false},
		{"4", []string{"no", "day", "week", "month", "year"}, "monthyear", false},
		{"5", []string{"no", "day", "week", "month", "year"}, " ", false},
		{"6", []string{"apple", "banana", "cherry"}, "apple", true},
		{"7", []string{"apple", "banana", "cherry"}, "banana", true},
		{"8", []string{"apple", "banana", "cherry"}, "pineapple", false},
		{"9", []string{"apple", "banana", "cherry"}, "APPLE", false},
		{"10", []string{"apple", "banana", "cherry"}, "", false},
		{"11", []string{}, "apple", false},
		{"12", []string{"apple", "banana", "cherry", " "}, " ", true},
		{"13", []string{" day", "week", "month"}, "day", false},
		{"14", []string{"day", "week", "month"}, "week", true},
		{"15", []string{"day", "week", "month"}, "year", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItemExists(tt.slice, tt.item); got != tt.expect {
				t.Errorf("ItemExists() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestUniqueInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"1", []int{1, 2, 3, 2, 4, 3, 5, 1}, []int{1, 2, 3, 4, 5}},
		{"2", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"3", []int{1, 1, 1, 1, 1}, []int{1}},
		{"4", []int{}, []int{}},
		{"5", []int{1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unique(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("Unique() = %v, want %v", got, tt.expected)
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Unique() = %v, want %v", got, tt.expected)
					return
				}
			}
		})
	}
}

func TestUniqueString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"1", []string{"apple", "banana", "apple", "cherry", "banana"}, []string{"apple", "banana", "cherry"}},
		{"2", []string{"apple", "banana", "cherry"}, []string{"apple", "banana", "cherry"}},
		{"3", []string{"apple", "apple", "apple"}, []string{"apple"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unique(tt.input)
			if len(got) != len(tt.expected) {
				t.Errorf("Unique() = %v, want %v", got, tt.expected)
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Unique() = %v, want %v", got, tt.expected)
					return
				}
			}
		})
	}
}

func TestEncodeBase64(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"hello"}, "aGVsbG8="},
		{"2", args{"world"}, "d29ybGQ="},
		{"3", args{"golang"}, "Z29sYW5n"},
		{"4", args{"123456"}, "MTIzNDU2"},
		{"5", args{"base64 encoding"}, "YmFzZTY0IGVuY29kaW5n"},
		{"6", args{"Привет, мир!"}, "0J/RgNC40LLQtdGCLCDQvNC40YAh"},
		{"7", args{"Hello World!"}, "SGVsbG8gV29ybGQh"},
		{"8", args{"Special! @#^&*()"}, "U3BlY2lhbCEgQCNeJiooKQ=="},
		{"9", args{"1234567890"}, "MTIzNDU2Nzg5MA=="},
		{"10", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeBase64(tt.args.text); got != tt.want {
				t.Errorf("EncodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeBase64(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"aGVsbG8="}, "hello"},
		{"2", args{"d29ybGQ="}, "world"},
		{"3", args{"Z29sYW5n"}, "golang"},
		{"4", args{"MTIzNDU2"}, "123456"},
		{"5", args{"YmFzZTY0IGVuY29kaW5n"}, "base64 encoding"},
		{"6", args{"0J/RgNC40LLQtdGCLCDQvNC40YAh"}, "Привет, мир!"},
		{"7", args{"SGVsbG8gV29ybGQh"}, "Hello World!"},
		{"8", args{"U3BlY2lhbCEgQCNeJiooKQ=="}, "Special! @#^&*()"},
		{"9", args{"MTIzNDU2Nzg5MA=="}, "1234567890"},
		{"10", args{""}, ""},
		{"11", args{"invalid_base64"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeBase64(tt.args.text); got != tt.want {
				t.Errorf("DecodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
