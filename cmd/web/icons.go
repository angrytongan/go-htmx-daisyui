package main

import (
	"net/http"
)

func (app *Application) icons(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{}

	app.render(w, r, "icons", pageData, http.StatusOK)
}
