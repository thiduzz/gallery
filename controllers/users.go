package controllers

import (
	"fmt"
	"github.com/thiduzz/lenslocked.com/views"
	"net/http"
)

type Users struct {
	IndexView *views.View
	NewView *views.View
}

func NewUsers() *Users {
	return &Users{
		IndexView: views.NewView("master", "views/users/index.gohtml"),
		NewView: views.NewView("master", "views/users/new.gohtml"),
	}
}

func (c *Users) Index(w http.ResponseWriter, r *http.Request)  {
	c.IndexView.Render(w, nil)
}

func (c *Users) New(w http.ResponseWriter, r *http.Request)  {
	c.NewView.Render(w, nil)
}

func (c Users) Store(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Created successfully")
}