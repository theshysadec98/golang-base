package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var pathToTemplate = "./templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	//User *data.User
}

func (app *Config) render(responseWriter http.ResponseWriter, request *http.Request, temp string, templateData *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", pathToTemplate),
		fmt.Sprintf("%s/header.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/navbar.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/footer.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/alerts.partial.gohtml", pathToTemplate),
	}

	var templateSlice []string

	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", pathToTemplate, temp))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	if templateData == nil {
		templateData = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		app.ErrorLog.Println(err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(responseWriter, app.AddDefaultData(templateData, request)); err != nil {
		app.ErrorLog.Println(err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Config) AddDefaultData(templateData *TemplateData, request *http.Request) *TemplateData {
	templateData.Flash = app.Session.PopString(request.Context(), "flash")
	templateData.Warning = app.Session.PopString(request.Context(), "warning")
	templateData.Error = app.Session.PopString(request.Context(), "error")

	if app.IsAuthenticated(request) {
		templateData.Authenticated = true
	}

	templateData.Now = time.Now()

	return templateData
}

func (app *Config) IsAuthenticated(request *http.Request) bool {
	return app.Session.Exists(request.Context(), "userID")
}
