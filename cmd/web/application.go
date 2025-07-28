package main

import (
	"bytes"
	"fmt"
	"ghdui/internal/kanban"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	initialColour = "red"
)

type Application struct {
	tpl               *template.Template
	darkMode          bool
	mapboxAccessToken string
	workflow          *kanban.Workflow
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
	workflow := kanban.NewMemory()

	// FIXME while we're testing, let's seed the kanban.
	workflow.AddNote(1, "one")
	workflow.AddNote(1, "two")
	workflow.AddNote(1, "three")
	workflow.AddNote(1, "four")
	workflow.AddNote(1, "five")

	workflow.AddNote(2, "six")
	workflow.AddNote(2, "seven")
	workflow.AddNote(2, "eight")
	workflow.AddNote(2, "nine")
	workflow.AddNote(2, "ten")

	workflow.AddNote(3, "eleven")
	workflow.AddNote(4, "twelve")
	workflow.AddNote(3, "thirteen")
	workflow.AddNote(4, "fourteen")
	workflow.AddNote(3, "fifteen")

	return &Application{
		tpl:               tpl,
		darkMode:          false,
		mapboxAccessToken: mustGetenv("MAPBOX_ACCESS_TOKEN"),
		workflow:          workflow,
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
		block += "-page"

		// Setup non-specific page template data here.
		if pageData == nil {
			pageData = map[string]any{}
		}

		pageData["NavLinks"] = []struct {
			Label string
			Href  string
		}{
			{Label: "root", Href: "/"},
			{Label: "fragments", Href: "/template-fragments"},
			{Label: "icons", Href: "/icons"},
			{Label: "badges", Href: "/badges"},
			{Label: "kanban", Href: "/kanban"},
		}

		pageData["DarkMode"] = app.darkMode
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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)

	_, err = w.Write(b.Bytes())
	if err != nil {
		app.serverError(w, r, fmt.Errorf("w.Write(): %w", err), http.StatusInternalServerError)

		return
	}
}
