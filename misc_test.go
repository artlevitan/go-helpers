// Copyright 2023-2025, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.2.2

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

func TestFileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"misc.go"}, true},
		{"2", args{"404.txt"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.args.filename); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
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
		s string
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
			if got := EncodeBase64(tt.args.s); got != tt.want {
				t.Errorf("EncodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeBase64(t *testing.T) {
	type args struct {
		s string
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
			if got := DecodeBase64(tt.args.s); got != tt.want {
				t.Errorf("DecodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActiveEnum(t *testing.T) {
	type args struct {
		flag string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1"}, "1"},
		{"2", args{"0"}, "0"},
		{"3", args{"2"}, "0"},
		{"4", args{""}, "0"},
		{"5", args{"true"}, "0"},
		{"6", args{"false"}, "0"},
		{"7", args{"yes"}, "0"},
		{"8", args{"no"}, "0"},
		{"9", args{"-1"}, "0"},
		{"10", args{"1.0"}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ActiveEnum(tt.args.flag); got != tt.want {
				t.Errorf("ActiveEnum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"1", "password123", false},
		{"2", "anotherPassword", false},
		{"3", "12345", false},
		{"4", "", false},
		{"5", "aVeryLongPasswordWithDifferentCharacters123!@#", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestComparePasswords(t *testing.T) {
	tests := []struct {
		name           string
		hashedPassword string
		password       string
		wantResult     bool
	}{
		{"1", "$2a$10$Eiknevkvo3H35yKPLC6z9eczlYQPSD5jE3kAALdkt6hi0DctTOp7O", "password123", true},
		{"2", "$2a$10$R7hVz7vGahLXxYOLuDKDCewBXl5kpDOjSPZgH6J/gPA6MzvSb/Pey", "wrongPassword", false},
		{"3", "$2a$10$AZIFEgFHsJILtwrmzb2tYufNCvTdKQLKMI0zmgdimh7w8njUExrXi", "12345", true},
		{"4", "$2a$10$Wn4dNE3jcYPVmnHg.kPs7uBtUJAVOfICkId81k7j/q4A0GZKrfVpS", "54321", false},
		{"5", "$2a$10$KzQHqmxnBeOvCwofE/dsyO9X/H1lANkGfdMN5XbfH7DZ4E4zp8n2u", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparePasswords(tt.hashedPassword, tt.password); got != tt.wantResult {
				t.Errorf("ComparePasswords() = %v, want %v", got, tt.wantResult)
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
		in any
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
		{"25", args{map[string]any{"a": 1, "b": "two", "c": 3.0}}, `{"a":1,"b":"two","c":3}`},
		{"26", args{map[string]any{"outer": map[string]any{"inner": "value"}}}, `{"outer":{"inner":"value"}}`},
		{"27", args{[3]int{1, 2, 3}}, "[1,2,3]"},
		{"28", args{func() *int { i := 42; return &i }()}, "42"},
		{"29", args{func() *Person { return &Person{Name: "Alice", Age: 25} }()}, `{"Name":"Alice","Age":25}`},
		{"30", args{complex(1, 2)}, "error:json: unsupported type: complex128"},
		{"31", args{[][]int{{1, 2}, {3, 4}}}, "[[1,2],[3,4]]"},
		{"32", args{[]Person{{Name: "Bob", Age: 30}, {Name: "Charlie", Age: 25}}}, `[{"Name":"Bob","Age":30},{"Name":"Charlie","Age":25}]`},
		{"33", args{any(Person{Name: "Diana", Age: 28})}, `{"Name":"Diana","Age":28}`},
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
