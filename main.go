package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	Count        int
	RedirectFrom string
	RedirectTo   string
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./redirect.html"))

	page := Page{
		Count:        5,
		RedirectFrom: "redirect01.hypermkt.jp",
		RedirectTo:   "http://www.yahoo.co.jp",
	}
	err := t.Execute(w, page)
	if err != nil {
		panic(err)
	}
}
