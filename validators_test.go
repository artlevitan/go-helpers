// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.12

package helpers

import (
	"testing"
)

func TestIsSQLDate(t *testing.T) {
	// Валидные даты
	validDates := []string{
		"2023-01-01",
		"1999-12-31",
		"2024-02-29", // високосный год
		"2020-06-15",
	}

	// Невалидные даты
	invalidDates := []string{
		"2023-13-01", // неверный месяц
		"2023-00-10", // неверный месяц
		"2023-12-32", // неверный день
		"2023-02-30", // неверный день
		"2001-02-29", // не високосный год
		"2023/12/01", // неправильный формат
		"23-12-01",   // недостаточно цифр в годе
	}

	for _, date := range validDates {
		if !IsSQLDate(date) {
			t.Errorf("Ожидалось, что дата %s будет валидной", date)
		}
	}

	for _, date := range invalidDates {
		if IsSQLDate(date) {
			t.Errorf("Ожидалось, что дата %s будет невалидной", date)
		}
	}
}

func TestIsSQLDateTime(t *testing.T) {
	// Валидные даты-время
	validDateTimes := []string{
		"2023-01-01 12:30:45",
		"1999-12-31 23:59:59",
		"2020-02-29 00:00:00", // високосный год
		"2022-08-15 15:45:30",
	}

	// Невалидные даты-время
	invalidDateTimes := []string{
		"2023-01-01 24:00:00", // неверное время
		"2023-12-01 23:60:00", // неверная минута
		"2023-12-01 12:30:60", // неверная секунда
		"2023-13-01 12:30:45", // неверный месяц
		"2023-12-32 12:30:45", // неверный день
		"2023-12-01T12:30:45", // неправильный формат
		"2023-01-01 12:30",    // недостаточно времени
	}

	for _, dateTime := range validDateTimes {
		if !IsSQLDateTime(dateTime) {
			t.Errorf("Ожидалось, что дата-время %s будет валидной", dateTime)
		}
	}

	for _, dateTime := range invalidDateTimes {
		if IsSQLDateTime(dateTime) {
			t.Errorf("Ожидалось, что дата-время %s будет невалидной", dateTime)
		}
	}
}

func TestIsSQLTime(t *testing.T) {
	// Валидное время
	validTimes := []string{
		"00:00:00",
		"23:59:59",
		"12:30:45",
		"07:45:00",
	}

	// Невалидное время
	invalidTimes := []string{
		"24:00:00",   // неверный час
		"12:60:00",   // неверная минута
		"12:00:60",   // неверная секунда
		"12:00",      // недостаточно времени
		"12:00:00:1", // лишняя секция
		"120000",     // отсутствуют разделители
	}

	for _, timeStr := range validTimes {
		if !IsSQLTime(timeStr) {
			t.Errorf("Ожидалось, что время %s будет валидным", timeStr)
		}
	}

	for _, timeStr := range invalidTimes {
		if IsSQLTime(timeStr) {
			t.Errorf("Ожидалось, что время %s будет невалидным", timeStr)
		}
	}
}

func TestIsHexColor(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"1f1f1F"}, true},
		{"2", args{"AFAFAF"}, true},
		{"3", args{"1AFFa1"}, true},
		{"4", args{"222fff"}, true},
		{"5", args{"F00"}, true},        // Ожидается true, так как "F00" — это валидный короткий HEX-код
		{"6", args{"#1f1f1F"}, true},    // HEX-код с символом #
		{"7", args{"#AFAFAF"}, true},    // HEX-код с символом #
		{"8", args{"#F00"}, true},       // Трехзначный HEX-код с символом #
		{"9", args{"GGGGGG"}, false},    // Неправильные символы
		{"10", args{"12345"}, false},    // HEX-код с неверной длиной
		{"11", args{""}, false},         // Пустая строка
		{"12", args{"1F1F1F1F"}, false}, // HEX-код слишком длинный
		{"13", args{"ZZZZZZ"}, false},   // Неверные символы
		{"14", args{"#123"}, true},      // Короткий HEX-код с символом #
		{"15", args{"FFF"}, true},       // Короткий HEX-код без символа #
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHexColor(tt.args.color); got != tt.want {
				t.Errorf("IsHexColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsURL(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{s: "http://appercase.ru"}, true},
		{"2", args{s: "https://appercase.ru"}, true},
		{"3", args{s: "ftp://appercase.ru"}, true},
		{"4", args{s: "http//appercase.ru"}, false},
		{"5", args{s: "appercase.ru"}, false},
		{"6", args{s: "https://www.example.com/path?query=123"}, true},
		{"7", args{s: "http://localhost:8080"}, true},
		{"8", args{s: "http://192.168.1.1"}, true},
		{"9", args{s: "ftp://192.168.1.1"}, true},
		{"10", args{s: "https://[::1]"}, true},
		{"11", args{s: "http://test-domain"}, true},
		{"12", args{s: "https://example.com:invalid"}, false},
		{"13", args{s: ""}, false},
		{"14", args{s: "https://"}, false},
		{"15", args{s: "https://www.example"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsURL(tt.args.s); got != tt.want {
				t.Errorf("IsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPv4(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{ip: "192.168.0.1"}, want: true},
		{name: "2", args: args{ip: "255.255.255.255"}, want: true},
		{name: "3", args: args{ip: "0.0.0.0"}, want: true},
		{name: "4", args: args{ip: "127.0.0.1"}, want: true},
		{name: "5", args: args{ip: "256.256.256.256"}, want: false},
		{name: "6", args: args{ip: "192.168.0"}, want: false},
		{name: "7", args: args{ip: "192.168.0.1.1"}, want: false},
		{name: "8", args: args{ip: "abc.def.ghi.jkl"}, want: false},
		{name: "9", args: args{ip: "::1"}, want: false},
		{name: "10", args: args{ip: "192.168.0.256"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIPv4(tt.args.ip); got != tt.want {
				t.Errorf("IsIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPv6(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{ip: "::1"}, want: true},
		{name: "2", args: args{ip: "2001:db8::ff00:42:8329"}, want: true},
		{name: "3", args: args{ip: "fe80::1ff:fe23:4567:890a"}, want: true},
		{name: "4", args: args{ip: "::"}, want: true},
		{name: "5", args: args{ip: "1200::AB00:1234::2552:7777:1313"}, want: false},
		{name: "6", args: args{ip: "1200::AB00:1234:O000:2552:7777:1313"}, want: false},
		{name: "7", args: args{ip: "192.168.0.1"}, want: false},
		{name: "8", args: args{ip: "2001:db8:85a3::8a2e:370:7334"}, want: true},
		{name: "9", args: args{ip: "12345::"}, want: false},
		{name: "10", args: args{ip: "1::1:1:1:1:1:1"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIPv6(tt.args.ip); got != tt.want {
				t.Errorf("IsIPv6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsJSON(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"1", args{`{}`}, true, false},
		{"2", args{`{"bar":2,"baz":3,"foo":1}`}, true, false},
		{"3", args{`{bar":2,"baz":3,"foo":1}`}, false, true},
		{"4", args{`{"bar":2,"baz":3,"foo":1,}`}, false, true},
		{"5", args{`"bar":2,"baz":3,"foo":1`}, false, true},
		{"6", args{`[]`}, true, false},
		{"7", args{`{"name":"John", "age":30, "city":"New York"}`}, true, false},
		{"8", args{`{"numbers":[1,2,3,4,5]}`}, true, false},
		{"9", args{`{"key":"value","list":[1,2,3],"obj":{"innerKey":"innerValue"}}`}, true, false},
		{"10", args{`{"unclosed_object":{"key":"value"`}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsJSON(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
