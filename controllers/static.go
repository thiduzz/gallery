package controllers

import (
	"fmt"
	"github.com/thiduzz/lenslocked.com/views"
	"net/http"
)

type Static struct {
	Home *views.View
	Contact *views.View
}

func NewStatic() *Static {
	return &Static{
		Home: views.NewView("master", "static/home"),
		Contact: views.NewView("master", "static/contact"),
	}
}

func (s *Static) NotFound(w http.ResponseWriter, r *http.Request)   {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404</h1>")
}