package main

import (
	"ghdui/internal/nav"
	"net/http"
)

func (app *Application) nav(w http.ResponseWriter, r *http.Request) {
	href := "/" + r.PathValue("page")
	queryParams := r.URL.RawQuery

	pageData := map[string]any{
		"Nav":         nav.MakeLinks(href),
		"Href":        href,
		"DarkMode":    app.darkMode,
		"QueryParams": queryParams,
	}

	app.render(w, r, "nav", pageData, http.StatusOK)
}
