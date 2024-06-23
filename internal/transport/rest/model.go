package rest

import "time"

type UrlToShort struct {
	Url    string `json:"url"`
	Author string `json:"author"`
}

type ShortenUrl struct {
	Original string    `json:"original"`
	Shorten  string    `json:"shorten"`
	Author   string    `json:"author"`
	Created  time.Time `json:"created"`
}
