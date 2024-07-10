package dtos

import (
	"time"
	"url-shortener/daos"
)

type UrlToShort struct {
	Url        string    `json:"url"`
	Expiration time.Time `json:"exp"`
	Author     string    `json:"author"`
}

type ShortenUrl struct {
	Original   string    `json:"original"`
	Shorten    string    `json:"shorten"`
	Author     User      `json:"author"`
	Created    time.Time `json:"created"`
	Expiration time.Time `json:"exp"`
}

func UrlDataToTransportModel(url daos.Url, user daos.User) *ShortenUrl {
	return &ShortenUrl{
		Original:   url.Url,
		Shorten:    url.Id,
		Expiration: url.Expiration,
		Created:    url.Created,
		Author:     User{Name: user.Name, Created: user.Created, Active: user.Active},
	}
}
