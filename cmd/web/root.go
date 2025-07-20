package main

import "net/http"

func (app *Application) root(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"NumChunks": 1,
	}

	app.render(w, r, "root", pageData, http.StatusOK)
}
