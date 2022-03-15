package httpserver

import "github.com/tobiassjosten/go-intro/pkg"

type User struct {
	Name string `json:"totallyrandomname"`
}

func NewDBUser(u *pkg.User) *User {
	return &User{
		Name: u.Name,
	}
}

func (dbu *User) User() *pkg.User {
	return &pkg.User{
		Name: dbu.Name,
	}
}
