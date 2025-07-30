package main

import (
	"ghdui/internal/people"
	"net/http"
)

type People struct {
	Name     string
	Age      int
	Archived bool
}

func (app *Application) templateFragments(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Value":  0,
		"People": people.All(),
	}

	app.render(w, r, "template-fragments", pageData, http.StatusOK)
}
