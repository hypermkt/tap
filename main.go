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

	// config := readConfig()
	// for _, redirect := range config.Redirects {
	// 	fmt.Printf("From: %s, To: %s", redirect.From, redirect.To)
	// }

	// TODO: リダイレクト元・先設定は別設定ファイルに移譲
	// TODO: ドメイン名をリダイレクト元が一致したときのみリダイレクトさせる
	// TODO: アクセスログ出力をする
	// TODO: ドメイン名が一致しなかった場合のエラー画面を出力する
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
