package main

import (
	"html/template"
	"log"
	"net/http"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("home.html", "templates/start.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Playhandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("play.html", "templates/tab.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Homehandler)
	http.HandleFunc("/play", Playhandler)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}
