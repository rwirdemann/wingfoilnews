package main

import (
	"bytes"
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

	simpleweb.Register("/new", func(w http.ResponseWriter, r *http.Request) {
		simpleweb.Render("templates/new.html", w, nil)
	}, "GET")

	simpleweb.Register("/publish", func(w http.ResponseWriter, r *http.Request) {
		title, err := simpleweb.FormValue(r, "title")
		if err != nil {
			simpleweb.Error(err.Error())
		}

		uri, err := simpleweb.FormValue(r, "url")
		if err != nil {
			simpleweb.Error(err.Error())
		}

		slog.Info("form values", "title", title, "url", uri)
		err = postNews(title, uri, "https://news.wingbuddies.de:8087/links")
		if err != nil {
			slog.Error("error", "error", err)
		}
		simpleweb.Redirect(w, r, "/")

	}, "POST")

	simpleweb.Run()
}

func postNews(title, uri, url string) error {
	link := wingfoilnews.Link{
		Title: title,
		URI:   uri,
	}
	bb, err := json.Marshal(link)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bb))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	return nil
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
