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

	// Widgets.
	mux.Get("/widget/server-time", app.widgetServerTime)

	// Widgets that we delay.
	mux.Group(func(r chi.Router) {
		r.Use(delayWidgets)

		r.Get("/widget/stat/{dataset}", app.widgetStat)
		r.Get("/widget/table/{dataset}", app.widgetTable)
		r.Get("/widget/timeline", app.widgetTimeline)
	})
}
