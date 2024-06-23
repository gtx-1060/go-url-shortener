package models

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	UserCreationError = errors.New("cant create user")
	UrlCreationError  = errors.New("cant create url")
	UrlGettingError   = errors.New("cant get url")
	UrlUpdatingError  = errors.New("cant update url")
	UserUpdatingError = errors.New("cant update user")
)

func CreateUser(db *sql.DB, user User) (uint64, error) {
	var id uint64
	row := db.QueryRow(
		"INSERT INTO m_user VALUES (NULL, ?, ?, ?) RETURNING id",
		user.name, user.created, user.active)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err)
		return 0, UserCreationError
	}
	return id, nil
}

func CreateUrl(db *sql.DB, url Url) error {
	_, err := db.Exec(
		"INSERT INTO url VALUES (?, ?, ?, ?, ?)",
		url.id, url.userId, url.url, url.created, url.active)
	if err != nil {
		fmt.Println(err)
		return UrlCreationError
	}
	return nil
}

func getUrl(db *sql.DB, id string) (string, error) {
	var url string
	row := db.QueryRow("SELECT url FROM url WHERE id == ?", id)
	if err := row.Scan(&url); err != nil {
		return url, err
	}
	return url, nil
}

func getUrlData(db *sql.DB, id string) (*Url, *User, error) {
	url := &Url{}
	user := &User{}
	row := db.QueryRow(`
		SELECT (url.id, url.url, url.created, url.active, mu.id, mu.name, mu.created, mu.active) 
		FROM url JOIN main.m_user mu on url.user_id = mu.id WHERE url.id == ?`, id)
	err := row.Scan(&url.id, &url.url, &url.created, &url.active, &url.id,
		&user.id, &user.name, &user.created, &user.active)
	if err != nil {
		fmt.Println(err)
		return nil, nil, UrlGettingError
	}
	url.userId = user.id
	return url, user, nil
}

func UpdateUrlAccessibility(db *sql.DB, id string, active bool) error {
	_, err := db.Exec("UPDATE url SET active=? WHERE id == ?", active, id)
	if err != nil {
		fmt.Println(err)
		return UrlUpdatingError
	}
	return nil
}

// UpdateUserAccessibility *Not Implemented*
func UpdateUserAccessibility(db *sql.DB, id uint64, active bool) error {
	_, err := db.Exec("UPDATE m_user SET active=? WHERE id == ?", active, id)
	if err != nil {
		fmt.Println(err)
		return UserUpdatingError
	}
	return nil
}

// RemoveUrl *Not Implemented*
func RemoveUrl(db *sql.DB, id string) error {
	// TODO
	return nil
}

func RemoveUser(db *sql.DB, id uint64) error {
	// TODO
	return nil
}
