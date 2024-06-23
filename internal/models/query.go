package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type DB interface {
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
}

var (
	UserCreationError = errors.New("cant create user")
	UrlCreationError  = errors.New("cant create Url")
	UrlGettingError   = errors.New("cant get Url")
	UrlUpdatingError  = errors.New("cant update Url")
	UserUpdatingError = errors.New("cant update user")
)

func CreateUser(db DB, user User) (uint64, error) {
	var id uint64
	row := db.QueryRow(
		"INSERT INTO m_user VALUES (NULL, ?, ?, ?) RETURNING Id",
		user.Name, user.Created, user.Active)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err)
		return 0, UserCreationError
	}
	return id, nil
}

func CreateUrl(db DB, url Url) error {
	_, err := db.Exec(
		"INSERT INTO Url VALUES (?, ?, ?, ?, ?)",
		url.Id, url.UserId, url.Url, url.Created, url.Active)
	if err != nil {
		fmt.Println(err)
		return UrlCreationError
	}
	return nil
}

func GetUser(db DB, name string) (*User, error) {
	user := &User{Name: name}
	row := db.QueryRow("SELECT Id, Active, Created FROM m_user WHERE Name == ?", name)
	if err := row.Scan(&user.Id, &user.Active, &user.Created); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUrl(db DB, id string) (string, error) {
	var url string
	row := db.QueryRow("SELECT Url FROM Url WHERE Id == ?", id)
	if err := row.Scan(&url); err != nil {
		return url, err
	}
	return url, nil
}

func GetUrlData(db DB, id string) (*Url, *User, error) {
	url := &Url{}
	user := &User{}
	row := db.QueryRow(`
		SELECT (Url.Id, Url.Url, Url.Created, Url.Active, mu.Id, mu.Name, mu.Created, mu.Active) 
		FROM Url JOIN main.m_user mu on Url.user_id = mu.Id WHERE Url.Id == ?`, id)
	err := row.Scan(&url.Id, &url.Url, &url.Created, &url.Active, &url.Id,
		&user.Id, &user.Name, &user.Created, &user.Active)
	if err != nil {
		fmt.Println(err)
		return nil, nil, UrlGettingError
	}
	url.UserId = user.Id
	return url, user, nil
}

func UpdateUrlAccessibility(db DB, id string, active bool) error {
	_, err := db.Exec("UPDATE Url SET Active=? WHERE Id == ?", active, id)
	if err != nil {
		fmt.Println(err)
		return UrlUpdatingError
	}
	return nil
}

// UpdateUserAccessibility *Not Implemented*
func UpdateUserAccessibility(db DB, id uint64, active bool) error {
	_, err := db.Exec("UPDATE m_user SET Active=? WHERE Id == ?", active, id)
	if err != nil {
		fmt.Println(err)
		return UserUpdatingError
	}
	return nil
}

// RemoveUrl *Not Implemented*
func RemoveUrl(db DB, id string) error {
	// TODO
	return nil
}

func RemoveUser(db DB, id uint64) error {
	// TODO
	return nil
}
