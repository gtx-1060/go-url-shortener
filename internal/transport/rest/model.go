package rest

import "time"

type UrlToShort struct {
	url    string
	author string
}

type ShortenUrl struct {
	original string
	shorten  string
	author   string
	created  time.Time
}
