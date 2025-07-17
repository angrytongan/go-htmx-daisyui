package main

import (
	"errors"
	"ghdui/internal/downloads"
	"ghdui/internal/employees"
	"ghdui/internal/pageviews"
	"net/http"
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
		stat = pageviews.Stat()

	case "downloads":
		stat = downloads.Stat()
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
		table = employees.Table()
	}

	pageData := map[string]any{
		"Table": table,
	}

	app.render(w, r, "widget-table", pageData, http.StatusOK)
}
