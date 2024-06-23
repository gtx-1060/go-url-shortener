package services

import "time"

type Error struct {
	Message string `json:"message"`
}

type UrlToShort struct {
	Url    string `json:"url"`
	Author string `json:"author"`
}

type User struct {
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Active  bool      `json:"active"`
}

type ShortenUrl struct {
	Original string    `json:"original"`
	Shorten  string    `json:"shorten"`
	Author   User      `json:"author"`
	Created  time.Time `json:"created"`
}

func ErrorResponse(msg string) Error {
	return Error{msg}
}
