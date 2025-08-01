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

func delayWidgets(ms int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Duration(rand.IntN(ms)) * time.Millisecond)

			next.ServeHTTP(w, r)
		})
	}
}

func (app *Application) setRoutes(mux *chi.Mux) {
	// Pages.
	mux.Get("/", app.root)
	mux.Get("/template-fragments", app.templateFragments)
	mux.Get("/icons", app.icons)
	mux.Get("/badges", app.badges)

	// Navigation.
	mux.Get("/nav/{page}", app.nav)
	mux.Get("/nav/", app.nav)

	// Widgets.
	mux.Get("/widget/server-time", app.widgetServerTime)
	mux.Get("/widget/button-only", app.widgetButtonOnly)
	mux.Post("/widget/radial-graphs-update", app.widgetRadialGraphsUpdate)
	mux.Post("/theme/toggle", app.themeToggle)

	// Widgets that we delay.
	mux.Group(func(r chi.Router) {
		r.Use(delayWidgets(maxMilliseconds))

		r.Get("/widget/stat/{dataset}", app.widgetStat)
		r.Get("/widget/table/{dataset}", app.widgetTable)
		r.Get("/widget/timeline", app.widgetTimeline)
		r.Get("/widget/graph-random", app.widgetGraphRandom)
		r.Get("/widget/leaflet", app.widgetLeaflet)
		r.Get("/widget/radial-graphs", app.widgetRadialGraphs)
		r.Get("/widget/mapbox", app.widgetMapbox)

		// Things we change.
		r.Patch("/people/archive-toggle", app.peopleArchiveToggle)
	})

	// Tabs.
	mux.Get("/widget/tabs/{name}", app.widgetTabs)
	mux.Get("/widget/tabs", app.widgetTabs)

}
