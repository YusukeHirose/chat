package main

import (
	"chat/data"
	"log"
	"net/http"
)

func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println("Parse is failed")
	}
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		log.Fatal(err, "Cannot create session")
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, _ := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/login", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}
