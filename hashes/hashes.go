// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/

package hashes

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// MD5Hash возвращает MD5-хеш строки.
// Хеш представляет собой строку длиной 32 символа в шестнадцатеричном формате.
func MD5Hash(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

// SHA1Hash возвращает SHA1-хеш строки.
// Хеш представляет собой строку длиной 40 символов в шестнадцатеричном формате.
func SHA1Hash(text string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(text)))
}

// SHA256Hash возвращает SHA256-хеш строки.
// Хеш представляет собой строку длиной 64 символа в шестнадцатеричном формате.
func SHA256Hash(text string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
}

// SHA512Hash возвращает SHA512-хеш строки.
// Хеш представляет собой строку длиной 128 символов в шестнадцатеричном формате.
func SHA512Hash(text string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(text)))
}
