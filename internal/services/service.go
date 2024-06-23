package services

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"url-shortener/internal/database"
	"url-shortener/internal/models"
	"url-shortener/internal/transport/rest"
)

func calcHash(url string) string {
	data := []byte(url)
	hash := md5.Sum(data)
	result := make([]rune, 0, 8)
	for _, c := range hash[:4] {
		result = append(result, rune(c))
	}
	for _, c := range hash[12:] {
		result = append(result, rune(c))
	}
	return string(result)
}

func urlToTransportModel(url models.Url, user models.User) *rest.ShortenUrl {
	return &rest.ShortenUrl{
		Original: url.Url,
		Shorten:  url.Id,
		Created:  url.Created,
		Author:   rest.User{Name: user.Name, Created: user.Created, Active: user.Active},
	}
}

func printIfError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func closeDbCon(db *sql.DB) {
	printIfError(db.Close())
}

func MakeShortUrl(url rest.UrlToShort) (*rest.ShortenUrl, error) {
	db := database.ConnectDB()
	defer closeDbCon(db)
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cant start a transaction")
	}

	urlModel := models.Url{Url: url.Url, Created: time.Now(), Active: true}
	userModel := models.User{Name: url.Author, Created: time.Now(), Active: true}

	if user, err := models.GetUser(tx, url.Author); err != nil {
		userId, err := models.CreateUser(tx, userModel)
		if err != nil {
			printIfError(tx.Rollback())
			return nil, err
		}
		userModel.Id = userId
	} else {
		userModel = *user
	}
	urlModel.UserId = userModel.Id

	for i := 0; i < 10; i++ {
		urlModel.Id = calcHash(url.Url)
		if err := models.CreateUrl(tx, urlModel); err == nil {
			if err := tx.Commit(); err != nil {
				return nil, errors.New("cant commit a transaction")
			}
			return urlToTransportModel(urlModel, userModel), nil
		}
	}
	printIfError(tx.Rollback())
	return nil, errors.New("cant create unique URL because of collisions or other stuff")
}

func GetUrlByShort(shortUrl string) (string, error) {
	db := database.ConnectDB()
	defer closeDbCon(db)
	return models.GetUrl(db, shortUrl)
}

func GetUrlDataByShort(shortUrl string) (*rest.ShortenUrl, error) {
	db := database.ConnectDB()
	defer closeDbCon(db)
	urlModel, userModel, err := models.GetUrlData(db, shortUrl)
	if err != nil {
		return nil, err
	}
	return urlToTransportModel(*urlModel, *userModel), nil
}
