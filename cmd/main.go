package main

import (
	"embed"
	"encoding/json"
	"github.com/rwirdemann/simpleweb"
	"github.com/rwirdemann/wingfoilnews"
	"log"
	"net/http"
	"time"
)

// Expects all HTML templates in $PROJECTROOT/templates
//
//go:embed all:templates
var templates embed.FS

func init() {
	// Required Init call to tell SimpleWeb about its embedded templates, list
	// of base templates (empty) and port
	simpleweb.Init(templates, []string{}, 3030)
}

func main() {
	simpleweb.Register("/", func(w http.ResponseWriter, r *http.Request) {
		var data = struct {
			Links []wingfoilnews.Link
		}{}
		err := getNews("https://news.wingbuddies.de:8087/links", &data)
		if err != nil {
			log.Fatal(err)
		}

		simpleweb.Render("templates/index.html", w, struct {
			Links []wingfoilnews.Link
		}{Links: data.Links})
	}, "GET")
	simpleweb.Run()
}

func getNews(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
