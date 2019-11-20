package main

import (
	"chat/data"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {

	//templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		_, err := session(writer, request)
		if err != nil {
			// TODO navbarをpublic用にする
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			// TODO navbarをprivate用にする
			generateHTML(writer, threads, "layout", "navbar", "index")
		}
	}
}
