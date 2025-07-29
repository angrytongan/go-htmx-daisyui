package main

import (
	"ghdui/internal/nav"
	"net/http"
)

func (app *Application) badges(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Nav": nav.MakeLinks(r.URL.String()),
	}

	app.render(w, r, "badges", pageData, http.StatusOK)
}
