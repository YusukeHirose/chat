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
}
