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

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./redirect.html"))
	notfound := template.Must(template.ParseFiles("./404.html"))

	config := readConfig()
	for _, redirect := range config.Redirects {
		if r.URL.Host == redirect.From {
			page := Page{
				Count:        5,
				RedirectFrom: redirect.From,
				RedirectTo:   redirect.To,
			}
			fmt.Printf("From: %s, To: %s", redirect.From, redirect.To)

			err := t.Execute(w, page)
			if err != nil {
				panic(err)
			}
			return
		}
	}

	// TODO: アクセスログ出力をする
	err := notfound.Execute(w, Page{})
	if err != nil {
		panic(err)
	}
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
