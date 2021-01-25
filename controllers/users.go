package controllers

import (
	"fmt"
	"github.com/thiduzz/lenslocked.com/views"
	"net/http"
)

type Users struct {
	IndexView *views.View
	CreateView *views.View
}

type Credential struct {
	Email string `schema:"email,required"`
	Password string `schema:"password,required"`
}

func NewUsers() *Users {
	return &Users{
		IndexView: views.NewView("master", "users/index"),
		CreateView: views.NewView("master", "users/create"),
	}
}

func (c *Users) Index(w http.ResponseWriter, r *http.Request)  {
	c.IndexView.Render(w, nil)
}

func (c *Users) Create(w http.ResponseWriter, r *http.Request)  {
	c.CreateView.Render(w, nil)
}

func (c Users) Store(w http.ResponseWriter, r *http.Request)  {
	var credential Credential
	if err := parseForm(r, &credential); err != nil{
		panic(err)
	}
	fmt.Fprintln(w, "Created successfully:", credential.Email, credential.Password)
}