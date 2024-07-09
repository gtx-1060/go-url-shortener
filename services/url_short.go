package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
	"log"
	"time"
	"url-shortener/daos"
	"url-shortener/dtos"
)

const (
	maxCreateUrlAttempts = 10
)

func (serv *Service) MakeShortUrl(ctx context.Context, url dtos.UrlToShort) (*dtos.ShortenUrl, error) {
	var shortenUrl *dtos.ShortenUrl
	//txOptions := sql.TxOptions{ReadOnly: false, Isolation: sql.}

	txError := serv.dao.StartTx(ctx, nil, func(query daos.RWQuery) error {
		urlModel := daos.Url{Url: url.Url, Created: time.Now(), Active: true}
		userModel := daos.User{Name: url.Author, Created: time.Now(), Active: true}

		user, err := query.GetUser(url.Author)

		if err != nil && errors.Is(err, sql.ErrNoRows) {
			// if user not found
			userModel.Id, err = query.CreateUser(userModel)
			if err != nil {
				return err
			}
		} else {
			// if user already exists
			userModel = *user
		}
		// set owner data to url entity
		urlModel.UserId = userModel.Id

		for i := 0; i < maxCreateUrlAttempts; i++ {
			urlModel.Id = generateShortVersion(url.Url)
			if err = query.CreateUrl(urlModel); err == nil {
				// on success
				shortenUrl = dtos.UrlDataToTransportModel(urlModel, userModel)
				return nil
			}
			if !errors.Is(err, sqlite3.ErrConstraint) {
				return err
			}
		}
		return UrlCollisionsError
	})
	if txError != nil {
		log.Printf("error while creating short version of URL: %v", txError)
		return nil, ShortUrlCreationError
	}
	return shortenUrl, nil
}

func (serv *Service) GetUrlByShort(shortUrl string) (string, error) {
	if r, err := serv.dao.GetUrl(shortUrl); err != nil {
		return r, errors.New("url not available")
	} else {
		return r, nil
	}
}

func (serv *Service) GetUrlDataByShort(shortUrl string) (*dtos.ShortenUrl, error) {
	urlModel, userModel, err := serv.dao.GetUrlData(shortUrl)
	if err != nil {
		return nil, err
	}
	return dtos.UrlDataToTransportModel(*urlModel, *userModel), nil
}
