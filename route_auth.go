package main

import (
	"chat/data"
	"log"
	"net/http"
)

func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

func signupAccount(writer http.ResponseWriter, request *http.Request) {
	log.Printf("info: signup account")
	err := request.ParseForm()
	if err != nil {
		log.Fatal(err, "error: Cannot parse form")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		log.Fatal(err, "error: Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.ExecuteTemplate(writer, "layout", nil)
}

func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Printf("error: Parse is failed")
	}
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		log.Fatal(err, "error: Cannot create session")
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
