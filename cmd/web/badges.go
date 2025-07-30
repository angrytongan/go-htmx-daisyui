package main

import (
	"net/http"
)

func (app *Application) badges(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{}

	app.render(w, r, "badges", pageData, http.StatusOK)
}
