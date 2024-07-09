package services

import (
	"errors"
	"url-shortener/daos"
)

var (
	ShortUrlCreationError = errors.New("cant create short url")
	UrlCollisionsError    = errors.New("cant create unique URL because of collisions")
)

type Service struct {
	dao *daos.Dao
}

func NewService(dao *daos.Dao) *Service {
	return &Service{dao}
}
