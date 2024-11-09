// Copyright 2023-2024, Appercase LLC. All rights reserved.
// https://www.appercase.ru/
//
// v1.1.15

package helpers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
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

// SHA3_224 возвращает SHA3-224 хеш строки.
// Строка длиной 56 символов в шестнадцатеричном формате.
func SHA3_224(s string) string {
	return fmt.Sprintf("%x", sha3.Sum224([]byte(s)))
}

// SHA3_256 возвращает SHA3-256 хеш строки.
// Строка длиной 64 символа в шестнадцатеричном формате.
func SHA3_256(s string) string {
	return fmt.Sprintf("%x", sha3.Sum256([]byte(s)))
}

// SHA3_384 возвращает SHA3-384 хеш строки.
// Строка длиной 96 символов в шестнадцатеричном формате.
func SHA3_384(s string) string {
	return fmt.Sprintf("%x", sha3.Sum384([]byte(s)))
}

// SHA3_512 возвращает SHA3-512 хеш строки.
// Строка длиной 128 символов в шестнадцатеричном формате.
func SHA3_512(s string) string {
	return fmt.Sprintf("%x", sha3.Sum512([]byte(s)))
}

// SHAKE128 возвращает SHAKE128 хеш строки с заданной длиной вывода в байтах.
func SHAKE128(s string, outputLength int) string {
	hash := make([]byte, outputLength)
	sha3.ShakeSum128(hash, []byte(s))
	return fmt.Sprintf("%x", hash)
}

// SHAKE256 возвращает SHAKE256 хеш строки с заданной длиной вывода в байтах.
func SHAKE256(s string, outputLength int) string {
	hash := make([]byte, outputLength)
	sha3.ShakeSum256(hash, []byte(s))
	return fmt.Sprintf("%x", hash)
}

// BLAKE2b_256 возвращает BLAKE2b-256 хеш строки.
// Строка длиной 64 символа в шестнадцатеричном формате.
func BLAKE2b_256(s string) string {
	hasher, _ := blake2b.New256(nil)
	hasher.Write([]byte(s))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// BLAKE2b_512 возвращает BLAKE2b-512 хеш строки.
// Строка длиной 128 символов в шестнадцатеричном формате.
func BLAKE2b_512(s string) string {
	hasher, _ := blake2b.New512(nil)
	hasher.Write([]byte(s))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// BLAKE2s_256 возвращает BLAKE2s-256 хеш строки.
// Строка длиной 64 символа в шестнадцатеричном формате.
func BLAKE2s_256(s string) string {
	hasher, _ := blake2s.New256(nil)
	hasher.Write([]byte(s))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
