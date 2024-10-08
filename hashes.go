// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.10

package helpers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// MD5 возвращает MD5-хеш строки.
// Строка длиной 32 символа в шестнадцатеричном формате.
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// SHA1 возвращает SHA1-хеш строки.
// Строка длиной 40 символов в шестнадцатеричном формате.
func SHA1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

// SHA256 возвращает SHA256-хеш строки.
// Строка длиной 64 символа в шестнадцатеричном формате.
func SHA256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

// SHA512 возвращает SHA512-хеш строки.
// Строка длиной 128 символов в шестнадцатеричном формате.
func SHA512(s string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
}

// SHA3_256 возвращает SHA3-256 хеш строки.
// Строка длиной 64 символа в шестнадцатеричном формате.
func SHA3_256(s string) string {
	return fmt.Sprintf("%x", sha3.Sum256([]byte(s)))
}

// SHA3_512 возвращает SHA3-512 хеш строки.
// Строка длиной 128 символов в шестнадцатеричном формате.
func SHA3_512(s string) string {
	return fmt.Sprintf("%x", sha3.Sum512([]byte(s)))
}
