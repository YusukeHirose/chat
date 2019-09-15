package data

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func Threads() (thread []Thread, err error) {
	return
}
