// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.8

package helpers

import (
	"testing"
)

func TestCutString(t *testing.T) {
	type args struct {
		s      string
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
			if got := CutString(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("CutString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearString(t *testing.T) {
	type args struct {
		s string
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
			if got := ClearString(tt.args.s); got != tt.want {
				t.Errorf("ClearString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearTextarea(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Line 1\n\nLine 2\n\nLine 3"}, "Line 1\nLine 2\nLine 3"},
		{"2", args{"Line 1\r\nLine 2\r\n\r\nLine 3"}, "Line 1\nLine 2\nLine 3"},
		{"3", args{"Line 1\n    \nLine 2"}, "Line 1\nLine 2"},
		{"4", args{"\n\n\nLine 1\n\n\nLine 2"}, "Line 1\nLine 2"},
		{"5", args{"Line 1\n\n\n\nLine 2"}, "Line 1\nLine 2"},
		{"6", args{"Line 1\n \n \nLine 2"}, "Line 1\nLine 2"},
		{"7", args{"\n\nLine 1\n\nLine 2\n\n"}, "Line 1\nLine 2"},
		{"8", args{"\r\nLine 1\r\nLine 2\r\n"}, "Line 1\nLine 2"},
		{"9", args{"Line 1\nLine 2\nLine 3"}, "Line 1\nLine 2\nLine 3"},
		{"10", args{"\n \n\n \nLine 1\n \nLine 2\n \n"}, "Line 1\nLine 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClearTextarea(tt.args.s); got != tt.want {
				t.Errorf("ClearTextarea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterLetters(t *testing.T) {
	type args struct {
		s string
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
		{"10", args{"Ёлки зелёные"}, "Ёлкизелёные"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterLetters(tt.args.s); got != tt.want {
				t.Errorf("FilterLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterDigits(t *testing.T) {
	type args struct {
		s string
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
			if got := FilterDigits(tt.args.s); got != tt.want {
				t.Errorf("FilterDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckStringLength(t *testing.T) {
	type args struct {
		str       string
		minLength int
		maxLength int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"abc", 1, 5}, true},
		{"2", args{"abc", 4, 5}, false},
		{"3", args{"abcdef", 3, 6}, true},
		{"4", args{"abcdef", 7, 10}, false},
		{"5", args{"привет", 3, 6}, true},
		{"6", args{"привет", 7, 10}, false},
		{"7", args{"", 0, 0}, true},
		{"8", args{"a", 0, 1}, true},
		{"9", args{"a", 2, 5}, false},
		{"10", args{"𐍈", 1, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckStringLength(tt.args.str, tt.args.minLength, tt.args.maxLength); got != tt.want {
				t.Errorf("CheckStringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveEmojis(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Hello, World!"}, "Hello, World!"},
		{"2", args{"Hello, 😃 World!"}, "Hello, World!"},
		{"3", args{"Emojis 😃🚀🌍 in the middle"}, "Emojis in the middle"},
		{"4", args{"No emojis here"}, "No emojis here"},
		{"5", args{"😃 Only emojis 😃"}, "Only emojis"},
		{"6", args{"Multiple emojis 😃😃😃 in a row"}, "Multiple emojis in a row"},
		{"7", args{"Emojis at the end 😃"}, "Emojis at the end"},
		{"8", args{"😃 Emojis at the start"}, "Emojis at the start"},
		{"9", args{""}, ""},
		{"10", args{"😃"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveEmojis(tt.args.s); got != tt.want {
				t.Errorf("RemoveEmojis() = %v, want %v", got, tt.want)
			}
		})
	}
}
