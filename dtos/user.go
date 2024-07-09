package dtos

import "time"

type User struct {
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Active  bool      `json:"active"`
}
