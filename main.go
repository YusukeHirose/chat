package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/comail/colog"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func SetLogger() {
	colog.SetDefaultLevel(colog.LInfo)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
}

func main() {
	mux := http.NewServeMux()
	SetLogger()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signupAccount", signupAccount)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/process", process)

	log.SetOutput(os.Stdout)
	server.ListenAndServe()
}
