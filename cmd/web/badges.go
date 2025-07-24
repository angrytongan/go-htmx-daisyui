package main

import "net/http"

func (app *Application) badges(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "badges", nil, http.StatusOK)
}
