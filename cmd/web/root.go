package main

import "net/http"

func (app *Application) root(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"NumChunks": 5,
	}

	app.render(w, r, "root-page", pageData, http.StatusOK)
}
