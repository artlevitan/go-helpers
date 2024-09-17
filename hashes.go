// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.6

package helpers

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// MD5 возвращает MD5-хеш строки.
// Хеш представляет собой строку длиной 32 символа в шестнадцатеричном формате.
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// SHA1 возвращает SHA1-хеш строки.
// Хеш представляет собой строку длиной 40 символов в шестнадцатеричном формате.
func SHA1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

// SHA256 возвращает SHA3-256 хеш строки.
// Хеш представляет собой строку длиной 64 символа в шестнадцатеричном формате.
func SHA256(s string) string {
	return fmt.Sprintf("%x", sha3.Sum256([]byte(s)))
}

// SHA512 возвращает SHA3-512 хеш строки.
// Хеш представляет собой строку длиной 128 символов в шестнадцатеричном формате.
func SHA512(s string) string {
	return fmt.Sprintf("%x", sha3.Sum512([]byte(s)))
}
