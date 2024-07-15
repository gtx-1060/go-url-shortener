package services

import (
	"crypto/md5"
	rand2 "math/rand/v2"
)

var salt [8]byte

const alphabet = "!\"#$%&'()*+,-.0123456789:;<=>@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~"

func safeRune(b byte) rune {
	r := rune(33 + b%94)
	if r == '?' || r == '/' || r == '\\' {
		r++
	}
	return r
}

func RandomFromString(url string) string {
	data := []byte(url)
	hash := md5.Sum(append(data, salt[:]...))
	result := make([]rune, 0, 8)
	for _, c := range hash[:4] {
		result = append(result, safeRune(c))
	}
	for _, c := range hash[12:] {
		result = append(result, safeRune(c))
	}
	for i := 0; i < len(salt); i++ {
		salt[i] += (hash[i] - 10) * 3
	}
	return string(result)
}

func RandomString(length int) string {
	result := make([]uint8, 0, length)
	mmax := uint(len(alphabet))
	for i := 0; i < length; i++ {
		result = append(result, alphabet[rand2.UintN(mmax)])
	}
	return string(result)
}
