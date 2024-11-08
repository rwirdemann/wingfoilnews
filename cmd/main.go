package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/rwirdemann/simpleweb"
	"github.com/rwirdemann/wingfoilnews"
	"log/slog"
	"net/http"
	"time"
)

// Expects all HTML templates in $PROJECTROOT/templates
//
//go:embed all:templates
var templates embed.FS

func init() {
	// Required Init call to tell SimpleWeb about its embedded templates, list of base templates and port
	simpleweb.Init(templates, []string{"templates/layout.html"}, 3030)
}

func main() {
	simpleweb.Register("/", func(w http.ResponseWriter, r *http.Request) {
		var data = struct {
			Links []wingfoilnews.Link
		}{}
		err := getNews("https://news.wingbuddies.de:8087/links", &data)
		if err != nil {
			slog.Error("error", "error", err)
			simpleweb.Error(err.Error())
		}
		simpleweb.Render("templates/index.html", w, struct {
			Links []wingfoilnews.Link
		}{Links: data.Links})
	}, "GET")

	simpleweb.Register("/about", func(w http.ResponseWriter, r *http.Request) {
		simpleweb.Render("templates/about.html", w, nil)
	}, "GET")

	simpleweb.Run()
}

func getNews(url string, target any) error {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return fmt.Errorf("status code %d", r.StatusCode)
	}
	return json.NewDecoder(r.Body).Decode(target)
}
