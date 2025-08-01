package main

import (
	"net/http"
)

func (app *Application) root(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"NumChunks":   1,
		"QueryParams": r.URL.RawQuery,
	}

	app.render(w, r, "/", pageData, http.StatusOK)
}
