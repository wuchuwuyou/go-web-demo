package controller

import (
    "html/template"

    "github.com/gorilla/sessions"
)

var (
    homeController  home
    templates       map[string]*template.Template
    sessionName     string
    store           *sessions.CookieStore
    flashName       string
    pageLimit       int
)

func init() {
    templates = PopulateTemplates()
    store = sessions.NewCookieStore([]byte("something-very-secret"))
    sessionName = "MW_GO_DEMO"
    flashName = "go-flash"
    pageLimit = 5
}

// Startup func
func Startup() {
    homeController.registerRoutes()
}