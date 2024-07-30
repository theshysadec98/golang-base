package main

import "net/http"

func (app *Config) HomePage(responseWriter http.ResponseWriter, request *http.Request) {
	app.render(responseWriter, request, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(responseWriter http.ResponseWriter, request *http.Request) {
	app.render(responseWriter, request, "login.page.gohtml", nil)
}

func (app *Config) PostLogin(responseWriter http.ResponseWriter, request *http.Request) {
	app.render(responseWriter, request, "login.page.gohtml", nil)
}

func (app *Config) Logout(responseWriter http.ResponseWriter, request *http.Request) {
	app.render(responseWriter, request, "login.page.gohtml", nil)
}

func (app *Config) RegisterPage(responseWriter http.ResponseWriter, request *http.Request) {
	app.render(responseWriter, request, "register.page.gohtml", nil)
}

func (app *Config) PostRegister(responseWriter http.ResponseWriter, request *http.Request) {
	app.render(responseWriter, request, "register.page.gohtml", nil)
}

func (app *Config) ActivateAccount(responseWriter http.ResponseWriter, request *http.Request) {
	app.render(responseWriter, request, "login.page.gohtml", nil)
}
