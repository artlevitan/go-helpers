// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/

package strings

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
		{"Empty string", args{"", 10}, ""},
		{"Length zero", args{"Lorem ipsum", 0}, ""},
		{"Short text, short cut", args{"Lorem ipsum dolor sit amet.", 10}, "Lorem ipsu"},
		{"Short text, exact length", args{"Lorem ipsum", 11}, "Lorem ipsum"},
		{"Long text, no cut", args{"Lorem ipsum", 20}, "Lorem ipsum"},
		{"Russian text", args{"Привет, мир!", 7}, "Привет,"},
		{"Emoji short", args{"😊😃😄😁", 2}, "😊😃"},
		{"Emoji long", args{"😊😃😄😁", 3}, "😊😃😄"},
		{"Japanese text", args{"こんにちは世界", 5}, "こんにちは"},
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
		{"11", args{"NoExtraSpaces"}, "NoExtraSpaces"},
		{"12", args{"       "}, ""},
		{"13", args{"a   b   c"}, "a b c"},
		{"14", args{"Text with\tmultiple    spaces \t and\t tabs"}, "Text with multiple spaces and tabs"},
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
		{"Only letters", args{"HelloWorld"}, "HelloWorld"},
		{"Mixed letters and digits", args{"H3lloW0rld"}, "HlloWrld"},
		{"Letters with special chars", args{"H@e#l%l^o*W(o)r_l+d!"}, "HelloWorld"},
		{"Only digits", args{"1234567890"}, ""},
		{"Mixed digits and special chars", args{"123@#$456&*7890"}, ""},
		{"Mixed letters, digits, and special chars", args{"H3l@lo2 W0r^ld!"}, "HlloWrld"},
		{"Letters with spaces", args{"  Hello   World  "}, "HelloWorld"},
		{"Mixed letters and non-English chars", args{"Привет Hello"}, "ПриветHello"},
		{"Empty string", args{""}, ""},
		{"Non-English letters", args{"Привет Мир"}, "ПриветМир"},
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
		{"Only digits", args{"1234567890"}, "1234567890"},
		{"Mixed letters and digits", args{"H3lloW0rld123"}, "30123"},
		{"Digits with special chars", args{"12@34#56$78"}, "12345678"},
		{"Only letters", args{"HelloWorld"}, ""},
		{"Mixed letters and special chars", args{"H3l@l#oW*r&ld!"}, "3"},
		{"Mixed digits, letters, and special chars", args{"1H2e3l4l5o6W7o8r9l0d!"}, "1234567890"},
		{"Digits with spaces", args{" 123 456 7890 "}, "1234567890"},
		{"Empty string", args{""}, ""},
		{"Non-English letters with digits", args{"Привет123"}, "123"},
		{"Non-English letters", args{"Привет Мир"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterDigits(tt.args.text); got != tt.want {
				t.Errorf("FilterDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
