package main

import "net/http"

func (app *Application) kanban(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "kanban", nil, http.StatusOK)
}
