package controllers

import (
	"fmt"
	"github.com/thiduzz/lenslocked.com/models"
	"github.com/thiduzz/lenslocked.com/views"
	"net/http"
)

type Users struct {
	IndexView *views.View
	CreateView *views.View
	service	*models.UserService
}

type Credential struct {
	Name string `schema:"name,required"`
	Email string `schema:"email,required"`
	Password string `schema:"password,required"`
}

func NewUsers(us *models.UserService) *Users {
	return &Users{
		IndexView: views.NewView("master", "users/index"),
		CreateView: views.NewView("master", "users/create"),
		service: us,
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
	user := models.User{
		Name:  credential.Name,
		Email: credential.Email,
	}
	err := c.service.Create(&user)
	if err != nil{
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}
	fmt.Fprint(w, user)
}