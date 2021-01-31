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

type RegistrationForm struct {
	Name string `schema:"name"`
	Email string `schema:"email,required"`
	Password string `schema:"password,required"`
}

type LoginForm struct {
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

func (c *Users) Login(w http.ResponseWriter, r *http.Request) {
	var form LoginForm
	if err := parseForm(r, &form); err != nil{
		panic(err)
	}
	user, err := c.service.Authenticate(form.Email, form.Password)
	switch err {
	case models.ErrNotFound:
		fmt.Fprintln(w, "Invalid email address.")
	case nil:
		fmt.Fprint(w, user)
	case models.ErrInvalidPassword:
		fmt.Fprintln(w, "Invalid password.")
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c Users) Store(w http.ResponseWriter, r *http.Request)  {
	var form RegistrationForm
	if err := parseForm(r, &form); err != nil{
		panic(err)
	}
	user := models.User{
		Name:  form.Name,
		Email: form.Email,
		Password: form.Password,
	}
	err := c.service.Create(&user)
	if err != nil{
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}
	fmt.Fprint(w, user)
}