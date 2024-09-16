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

func TestCreateCacheKey(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{nil}, "nil"},
		{"2", args{1}, "1"},
		{"3", args{uint(0)}, "0"},
		{"4", args{uint8(255)}, "255"},
		{"5", args{uint16(65535)}, "65535"},
		{"6", args{uint32(4294967295)}, "4294967295"},
		{"7", args{uint64(18446744073709551615)}, "18446744073709551615"},
		{"8", args{"1"}, "1"},
		{"9", args{[]uint8{52, 56}}, "48"},
		{"10", args{[]int{1, 2, 3, 4, 5}}, "12345"},
		{"11", args{[]int8{-128, 0, 127}}, "-1280127"},
		{"12", args{[]int16{-32768, 0, 32767}}, "-32768032767"},
		{"13", args{[]int32{-2147483648, 0, 2147483647}}, "-214748364802147483647"},
		{"14", args{[]int64{-9223372036854775808, 0, 9223372036854775807}}, "-922337203685477580809223372036854775807"},
		{"15", args{[]uint{52, 56}}, "5256"},
		{"16", args{[]uint{0, 1, 4294967295}}, "014294967295"},
		{"17", args{[]uint16{52, 56}}, "5256"},
		{"18", args{[]uint32{52, 56}}, "5256"},
		{"19", args{[]uint64{52, 56}}, "5256"},
		{"20", args{[]string{"1", "2", "3", "4", "5"}}, `["1","2","3","4","5"]`},
		{"21", args{true}, "true"},
		{"22", args{false}, "false"},
		{"23", args{3.14159}, "3.14159"},
		{"24", args{[]bool{true, false, true}}, `[true,false,true]`},
		{"25", args{map[string]interface{}{"a": 1, "b": "two", "c": 3.0}}, `{"a":1,"b":"two","c":3}`},
		{"26", args{map[string]interface{}{"outer": map[string]interface{}{"inner": "value"}}}, `{"outer":{"inner":"value"}}`},
		{"27", args{[3]int{1, 2, 3}}, "[1,2,3]"},
		{"28", args{func() *int { i := 42; return &i }()}, "42"},
		{"29", args{func() *Person { return &Person{Name: "Alice", Age: 25} }()}, `{"Name":"Alice","Age":25}`},
		{"30", args{complex(1, 2)}, "error:json: unsupported type: complex128"},
		{"31", args{[][]int{{1, 2}, {3, 4}}}, "[[1,2],[3,4]]"},
		{"32", args{[]Person{{Name: "Bob", Age: 30}, {Name: "Charlie", Age: 25}}}, `[{"Name":"Bob","Age":30},{"Name":"Charlie","Age":25}]`},
		{"33", args{interface{}(Person{Name: "Diana", Age: 28})}, `{"Name":"Diana","Age":28}`},
		{"34", args{float32(2.71828)}, "2.71828"},
		{"35", args{[]float64{1.1, 2.2, 3.3}}, `[1.1,2.2,3.3]`},
	}
	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateCacheKey(tt.args.in); got != tt.want {
				t.Errorf("CreateCacheKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
