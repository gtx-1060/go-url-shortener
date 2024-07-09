package services

import "crypto/md5"

var salt [8]byte

func safeRune(b byte) rune {
	r := rune(33 + b%94)
	if r == '?' || r == '/' || r == '\\' {
		r++
	}
	return r
}

// TODO: Simplify code below
func generateShortVersion(url string) string {
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
