package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"text/template"

	_ "github.com/hypermkt/tap/statik"
	config "github.com/hypermkt/tap/tap"
	"github.com/rakyll/statik/fs"
)

type Page struct {
	Count        int
	RedirectFrom string
	RedirectTo   string
}

func main() {
	http.HandleFunc("/", handler)
	// serve static files
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(getAssetsFS())))
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
	for _, redirect := range config.ReadConfig().Redirects {
		if r.Host == parseURL(redirect.From).Host {
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

func parseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func getAssetsFS() http.FileSystem {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	return statikFS
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
