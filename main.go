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
