// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.2

package helpers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// MD5 возвращает MD5-хеш строки.
// Хеш представляет собой строку длиной 32 символа в шестнадцатеричном формате.
func MD5(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

// SHA1 возвращает SHA1-хеш строки.
// Хеш представляет собой строку длиной 40 символов в шестнадцатеричном формате.
func SHA1(text string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(text)))
}

// SHA256 возвращает SHA256-хеш строки.
// Хеш представляет собой строку длиной 64 символа в шестнадцатеричном формате.
func SHA256(text string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
}

// SHA512 возвращает SHA512-хеш строки.
// Хеш представляет собой строку длиной 128 символов в шестнадцатеричном формате.
func SHA512(text string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(text)))
}
