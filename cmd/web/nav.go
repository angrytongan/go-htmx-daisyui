package main

import (
	"fmt"
	"ghdui/internal/nav"
	"net/http"
)

func (app *Application) nav(w http.ResponseWriter, r *http.Request) {
	href := "/" + r.PathValue("page")

	fmt.Println("href", href)

	pageData := map[string]any{
		"Nav":  nav.MakeLinks(href),
		"Href": href,
	}

	fmt.Println("loading", href, pageData["Nav"])

	app.render(w, r, "nav", pageData, http.StatusOK)
}
