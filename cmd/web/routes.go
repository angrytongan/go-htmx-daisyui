package main

import (
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	maxMilliseconds = 2000
)

func delayWidgets(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(rand.IntN(maxMilliseconds)) * time.Millisecond)

		next.ServeHTTP(w, r)
	})
}

func (app *Application) setRoutes(mux *chi.Mux) {
	// Pages.
	mux.Get("/", app.root)
	mux.Get("/template-fragments", app.templateFragments)

	// Widgets.
	mux.Get("/widget/server-time", app.widgetServerTime)
	mux.Get("/widget/button-only", app.widgetButtonOnly)

	// Widgets that we delay.
	mux.Group(func(r chi.Router) {
		r.Use(delayWidgets)

		r.Get("/widget/stat/{dataset}", app.widgetStat)
		r.Get("/widget/table/{dataset}", app.widgetTable)
		r.Get("/widget/timeline", app.widgetTimeline)
		r.Get("/widget/graph-random", app.widgetGraphRandom)
		r.Get("/widget/map", app.widgetMap)

		// Things we change.
		r.Patch("/people/archive-toggle", app.peopleArchiveToggle)
	})
}
