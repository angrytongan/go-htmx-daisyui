package main

import (
	"errors"
	"ghdui/internal/downloads"
	"ghdui/internal/employees"
	"ghdui/internal/pageviews"
	"ghdui/internal/registers"
	"ghdui/internal/users"
	"net/http"
	"time"
)

var (
	ErrMissingDataset = errors.New("missing dataset")
)

func (app *Application) widgetStat(w http.ResponseWriter, r *http.Request) {
	dataset := r.PathValue("dataset")
	if dataset == "" {
		app.serverError(w, r, ErrMissingDataset, http.StatusBadRequest)

		return
	}

	var stat any

	switch dataset {
	case "pageviews":
		time.Sleep(500 * time.Millisecond)
		stat = pageviews.Stat()

	case "downloads":
		time.Sleep(2 * time.Second)
		stat = downloads.Stat()

	case "new-registers":
		time.Sleep(750 * time.Millisecond)
		stat = registers.Stat()

	case "users":
		time.Sleep(1500 * time.Millisecond)
		stat = users.Stat()
	}

	pageData := map[string]any{
		"Stat": stat,
	}

	app.render(w, r, "widget-stat", pageData, http.StatusOK)
}

func (app *Application) widgetTable(w http.ResponseWriter, r *http.Request) {
	dataset := r.PathValue("dataset")
	if dataset == "" {
		app.serverError(w, r, ErrMissingDataset, http.StatusBadRequest)

		return
	}

	var table any

	switch dataset {
	case "employees":
		time.Sleep(1100 * time.Millisecond)
		table = employees.Table()
	}

	pageData := map[string]any{
		"Table": table,
	}

	app.render(w, r, "widget-table", pageData, http.StatusOK)
}
