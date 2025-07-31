package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Application struct {
	tpl               *template.Template
	darkMode          bool
	mapboxAccessToken string
}

func mustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("missing %s", key)
	}

	return value
}

func newApplication() *Application {
	tpl := template.Must(template.ParseGlob("./templates/*.tmpl"))

	return &Application{
		tpl:               tpl,
		darkMode:          false,
		mapboxAccessToken: mustGetenv("MAPBOX_ACCESS_TOKEN"),
	}
}

func (app *Application) serverError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
	statusCode int,
) {
	log.Printf("%s %s: %v\n", r.Method, r.URL, err)
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func (app *Application) render(w http.ResponseWriter,
	r *http.Request,
	block string,
	pageData map[string]any,
	statusCode int,
) {
	var b bytes.Buffer

	// Render a full page if we didn't get a htmx request.
	if r.Header.Get("Hx-Request") != "true" {
		// Setup non-specific page template data here.
		if pageData == nil {
			pageData = map[string]any{}
		}

		// Setup links, set active page.
		pageData["Page"] = block
		if block == "/" {
			pageData["Page"] = ""
		}

		pageData["DarkMode"] = app.darkMode

		block += "-page"
	}

	err := app.tpl.ExecuteTemplate(&b, block, pageData)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("app.tpl.ExecuteTemplate(%s): %w", block, err),
			http.StatusInternalServerError,
		)

		return
	}

	// NOTE hardcoded content type here
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)

	_, err = w.Write(b.Bytes())
	if err != nil {
		app.serverError(w, r, fmt.Errorf("w.Write(): %w", err), http.StatusInternalServerError)

		return
	}
}
