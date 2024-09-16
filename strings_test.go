// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/

package helpers

import (
	"testing"
)

func TestCutString(t *testing.T) {
	type args struct {
		text   string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"", 10}, ""},
		{"2", args{"Lorem ipsum", 0}, ""},
		{"3", args{"Lorem ipsum dolor sit amet.", 10}, "Lorem ipsu"},
		{"4", args{"Lorem ipsum", 11}, "Lorem ipsum"},
		{"5", args{"Lorem ipsum", 20}, "Lorem ipsum"},
		{"6", args{"Привет, мир!", 7}, "Привет,"},
		{"7", args{"😊😃😄😁", 2}, "😊😃"},
		{"8", args{"😊😃😄😁", 3}, "😊😃😄"},
		{"9", args{"こんにちは世界", 5}, "こんにちは"},
		{"10", args{"Test long string", 4}, "Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutString(tt.args.text, tt.args.length); got != tt.want {
				t.Errorf("CutString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearString(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Hello     World"}, "Hello World"},
		{"2", args{"   Hello           "}, "Hello"},
		{"3", args{"  Hello         World  "}, "Hello World"},
		{"4", args{"  Привет!     Как дела?     "}, "Привет! Как дела?"},
		{"5", args{"  Какая   чистая    строка   "}, "Какая чистая строка"},
		{"6", args{"  "}, ""},
		{"7", args{"\tHello\tWorld\t"}, "Hello World"},
		{"8", args{"\nHello \n World\n"}, "Hello World"},
		{"9", args{"Hello\n\n\nWorld"}, "Hello World"},
		{"10", args{"Привет,   мир!"}, "Привет, мир!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearString(tt.args.text); got != tt.want {
				t.Errorf("ClearString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLetters(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"HelloWorld"}, "HelloWorld"},
		{"2", args{"H3lloW0rld"}, "HlloWrld"},
		{"3", args{"H@e#l%l^o*W(o)r_l+d!"}, "HelloWorld"},
		{"4", args{"1234567890"}, ""},
		{"5", args{"123@#$456&*7890"}, ""},
		{"6", args{"H3l@lo2 W0r^ld!"}, "HlloWrld"},
		{"7", args{"  Hello   World  "}, "HelloWorld"},
		{"8", args{"Привет Hello"}, "ПриветHello"},
		{"9", args{""}, ""},
		{"10", args{"Привет Мир"}, "ПриветМир"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterLetters(tt.args.text); got != tt.want {
				t.Errorf("FilterLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDigits(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1234567890"}, "1234567890"},
		{"2", args{"H3lloW0rld123"}, "30123"},
		{"3", args{"12@34#56$78"}, "12345678"},
		{"4", args{"HelloWorld"}, ""},
		{"5", args{"H3l@l#oW*r&ld!"}, "3"},
		{"6", args{"1H2e3l4l5o6W7o8r9l0d!"}, "1234567890"},
		{"7", args{" 123 456 7890 "}, "1234567890"},
		{"8", args{""}, ""},
		{"9", args{"Привет123"}, "123"},
		{"10", args{"Привет Мир"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterDigits(tt.args.text); got != tt.want {
				t.Errorf("FilterDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
