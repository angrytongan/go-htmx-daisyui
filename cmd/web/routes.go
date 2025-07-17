package main

import "github.com/go-chi/chi/v5"

func (app *Application) setRoutes(mux *chi.Mux) {
	// Pages.
	mux.Get("/", app.root)

	// Widgets.
	mux.Get("/widget/stat/{dataset}", app.widgetStat)
	mux.Get("/widget/table/{dataset}", app.widgetTable)
}
