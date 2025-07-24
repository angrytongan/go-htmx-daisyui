package main

import "net/http"

func (app *Application) themeToggle(w http.ResponseWriter, r *http.Request) {
	app.darkMode = !app.darkMode

	pageData := map[string]any{
		"DarkMode": app.darkMode,
	}

	app.render(w, r, "widget-theme-toggle-checkbox", pageData, http.StatusOK)
}
