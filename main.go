package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Count        int
	RedirectFrom string
	RedirectTo   string
}

type Config struct {
	Redirects []Redirect `json:"redirects"`
}

type Redirect struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func main() {
	http.HandleFunc("/", handler)
	// serve static files
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(getPort(), nil)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	return ":" + port
}

func handler(w http.ResponseWriter, r *http.Request) {
	for _, redirect := range readConfig().Redirects {
		if r.Host == redirect.From {
			displayRedirectPage(w, Page{
				Count:        5,
				RedirectFrom: redirect.From,
				RedirectTo:   redirect.To,
			})
			return
		}
	}

	displayNotFoundPage(w)
}

func readConfig() *Config {
	configJSON, err := ioutil.ReadFile("./config.json")
	jsonBytes := ([]byte)(configJSON)
	if err != nil {
		panic(err)
	}
	data := new(Config)
	err = json.Unmarshal(jsonBytes, data)
	if err != nil {
		fmt.Println("error:", err)
	}

	return data
}

func displayNotFoundPage(w http.ResponseWriter) {
	t := template.Must(template.ParseFiles("./templates/404.html"))
	w.WriteHeader(http.StatusNotFound)
	err := t.Execute(w, Page{})
	if err != nil {
		panic(err)
	}
}

func displayRedirectPage(w http.ResponseWriter, p Page) {
	w.WriteHeader(http.StatusOK)
	t := template.Must(template.ParseFiles("./templates/redirect.html"))
	err := t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}
