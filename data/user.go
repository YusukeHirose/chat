package data

import "time"

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	// TODO emailで検索
	return
}
