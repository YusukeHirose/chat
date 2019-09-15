package main

import (
	"chat/data"
	"net/http"
	"text/template"
)

func index(writer http.ResponseWriter, request *http.Request) {

	//templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		_, err := session(writer, request)
		public_tmpl_files := []string{"templates/layout.html", "templates/public_navbar.html", "templates/index.html"}
		private_tmpl_files := []string{"templates/layout.html", "templates/public_navbar.html", "templates/index.html"}
		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(private_tmpl_files...))
		} else {
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		}
		templates.ExecuteTemplate(writer, "layout", threads)
	}
}
