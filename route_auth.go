package main

import (
	"log"
	"net/http"
)

func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println("Parse is failed")
	}
}
