package main

import (
	"ghdui/internal/people"
	"net/http"
)

func (app *Application) peopleArchive(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")

	people.Archive(name)

	pageData := map[string]any{
		"Name":     name,
		"Archived": true,
	}

	app.render(w, r, "people-do-archive", pageData, http.StatusOK)
}

func (app *Application) peopleUnarchive(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")

	people.Unarchive(name)

	pageData := map[string]any{
		"Name":     name,
		"Archived": false,
	}

	app.render(w, r, "people-do-archive", pageData, http.StatusOK)
}
