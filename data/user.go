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

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    string
	CreatedAt time.Time
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	// TODO emailで検索
	return
}

func (user *User) CreateSession() (session Session, err error) {
	// TODO DB処理
	return
}

func (session *Session) Check() (valid bool, err error) {
	// TODO DB処理
	return
}
