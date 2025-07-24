package main

import "net/http"

func (app *Application) icons(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "icons", nil, http.StatusOK)
}
