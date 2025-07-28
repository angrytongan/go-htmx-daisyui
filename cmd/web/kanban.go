package main

import (
	"net/http"
)

func (app *Application) kanban(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Columns": app.workflow.Columns,
	}

	app.render(w, r, "kanban", pageData, http.StatusOK)
}

func (app *Application) kanbanProcess(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
