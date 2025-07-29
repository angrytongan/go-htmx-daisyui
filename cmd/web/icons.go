package main

import (
	"ghdui/internal/nav"
	"net/http"
)

func (app *Application) icons(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Nav": nav.MakeLinks(r.URL.String()),
	}

	app.render(w, r, "icons", pageData, http.StatusOK)
}
