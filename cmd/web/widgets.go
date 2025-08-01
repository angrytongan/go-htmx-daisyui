package main

import (
	"errors"
	"ghdui/internal/downloads"
	"ghdui/internal/employees"
	"ghdui/internal/pageviews"
	"ghdui/internal/registers"
	"ghdui/internal/users"
	"html/template"
	"net/http"
	"strconv"
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

func (app *Application) widgetButtonOnly(w http.ResponseWriter, r *http.Request) {
	value, _ := strconv.Atoi(r.URL.Query().Get("value"))

	pageData := map[string]any{
		"Value": value + 1,
	}

	app.render(w, r, "button-only", pageData, http.StatusOK)
}

func (app *Application) widgetLeaflet(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"ID": template.JS(strconv.FormatInt(time.Now().UnixMilli(), 10)),
	}

	app.render(w, r, "widget-leaflet", pageData, http.StatusOK)
}

func (app *Application) widgetMapbox(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"ID":                template.JS(strconv.FormatInt(time.Now().UnixMilli(), 10)),
		"MapboxAccessToken": app.mapboxAccessToken,
	}

	app.render(w, r, "widget-mapbox", pageData, http.StatusOK)
}

func (app *Application) widgetRadialGraphs(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Progress": 50,
	}

	app.render(w, r, "widget-radial-graphs", pageData, http.StatusOK)
}
func (app *Application) widgetRadialGraphsUpdate(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	pageData := map[string]any{
		"Progress": r.FormValue("val"),
	}

	app.render(w, r, "widget-radial-graphs-update", pageData, http.StatusOK)
}

func (app *Application) widgetTabs(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	params := r.URL.Query()
	tabs := []string{"one", "two", "three", "four", "five"}

	// If we were called with no argument, check to see if we have a query
	// parameter. Use that as the tab if we do.
	if name == "" {
		lastTab := params.Get("tab")
		if lastTab != "" {
			name = lastTab
		}
	}

	// If we still don't have a name, use the first tab.
	if name == "" {
		name = tabs[0]
	}

	pageData := map[string]any{
		"Tabs":      tabs,
		"ChosenTab": name,
		"Content":   "this is content for " + name,
	}

	app.render(w, r, "widget-tabs", pageData, http.StatusOK)
}
