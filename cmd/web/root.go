package main

import (
	"ghdui/internal/nav"
	"net/http"
)

func (app *Application) root(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Nav":       nav.MakeLinks(r.URL.String()),
		"NumChunks": 1,
	}

	app.render(w, r, "root", pageData, http.StatusOK)
}
