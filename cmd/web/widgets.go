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
		stat = pageviews.Stat()

	case "downloads":
		stat = downloads.Stat()

	case "new-registers":
		stat = registers.Stat()

	case "users":
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
		table = employees.Table()
	}

	pageData := map[string]any{
		"Table": table,
	}

	app.render(w, r, "widget-table", pageData, http.StatusOK)
}

func (app *Application) widgetServerTime(w http.ResponseWriter, r *http.Request) {
	layout := "2006-01-02 15:04"
	pageData := map[string]any{
		"Now": time.Now().Format(layout),
	}

	app.render(w, r, "widget-server-time", pageData, http.StatusOK)
}

func (app *Application) widgetTimeline(w http.ResponseWriter, r *http.Request) {
	items := map[string]string{
		"1984": "First Macintosh computer",
		"1998": "iMac",
		"2001": "iPod",
		"2007": "iPhone",
		"2015": "Apple Watch",
	}

	pageData := map[string]any{
		"Items": items,
	}

	app.render(w, r, "widget-timeline", pageData, http.StatusOK)
}
