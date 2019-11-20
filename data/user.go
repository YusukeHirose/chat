package data

import (
	"log"
	"time"

	"github.com/google/uuid"
)

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

aaaa

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

func (user *User) Create() (err error) {
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Fatal("err!!!!")
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(uuid.New(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	log.Print("use is created!")
	return
}
