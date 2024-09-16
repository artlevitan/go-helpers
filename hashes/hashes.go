// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/

package hashes

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"strconv"
	"time"
)

// GenerateMD5Hash возвращает MD5-хеш строки.
// Хеш представляет собой строку длиной 32 символа в шестнадцатеричном формате.
func GenerateMD5Hash(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

// GenerateSHA1Hash возвращает SHA1-хеш строки.
// Хеш представляет собой строку длиной 40 символов в шестнадцатеричном формате.
func GenerateSHA1Hash(text string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(text)))
}

// GenerateSHA256Hash возвращает SHA256-хеш строки.
// Хеш представляет собой строку длиной 64 символа в шестнадцатеричном формате.
func GenerateSHA256Hash(text string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
}

// GenerateSHA512Hash возвращает SHA512-хеш строки.
// Хеш представляет собой строку длиной 128 символов в шестнадцатеричном формате.
func GenerateSHA512Hash(text string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(text)))
}

// GenerateRandomMD5Hash генерирует случайный MD5-хеш.
func GenerateRandomMD5Hash() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return GenerateMD5Hash(timestamp)
}

// GenerateRandomSHA1Hash генерирует случайный SHA1-хеш.
func GenerateRandomSHA1Hash() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return GenerateSHA1Hash(timestamp)
}

// GenerateRandomSHA256Hash генерирует случайный SHA256-хеш.
func GenerateRandomSHA256Hash() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return GenerateSHA256Hash(timestamp)
}

// GenerateRandomSHA512Hash генерирует случайный SHA512-хеш.
func GenerateRandomSHA512Hash() string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return GenerateSHA512Hash(timestamp)
}
