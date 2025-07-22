package main

import (
	"ghdui/internal/people"
	"net/http"
)

func (app *Application) peopleArchiveToggle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	archived := people.ArchiveToggle(name)
	person := people.Find(name)

	pageData := map[string]any{
		"Name":     person.Name,
		"Age":      person.Age,
		"Archived": archived,
	}

	app.render(w, r, "people-do-archive", pageData, http.StatusOK)
}
