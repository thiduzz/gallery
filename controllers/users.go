package controllers

import (
	"github.com/thiduzz/lenslocked.com/helpers"
	"github.com/thiduzz/lenslocked.com/models"
	"github.com/thiduzz/lenslocked.com/rand"
	"github.com/thiduzz/lenslocked.com/views"
	"log"
	"net/http"
)

type Users struct {
	IndexView *views.View
	CreateView *views.View
	ShowView *views.View
	service	models.UserService
}

type RegistrationForm struct {
	Name string `schema:"name"`
	Email string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers(us models.UserService) *Users {
	return &Users{
		IndexView: views.NewView("master", "users/index"),
		CreateView: views.NewView("master", "users/create"),
		ShowView: views.NewView("master", "users/show"),
		service: us,
	}
}

func (c *Users) Index(w http.ResponseWriter, r *http.Request)  {
	c.IndexView.Render(w, r,nil)
}

func (c *Users) Create(w http.ResponseWriter, r *http.Request)  {
	c.CreateView.Render(w, r,nil)
}

func (c *Users) Show(w http.ResponseWriter, r *http.Request)  {
	email, err := r.Cookie("remember_token")
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	c.ShowView.Render(w,r,email)
}

func (c *Users) Login(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form LoginForm
	if err := helpers.ParseForm(r, &form); err != nil{
		log.Print(err)
		vd.SetAlert(err)
		c.IndexView.Render(w,r,vd)
		return
	}
	user, err := c.service.Authenticate(form.Email, form.Password)
	if err != nil{
		switch err {
		case models.ErrNotFound:
			vd.SetAlertError("Invalid email address")
		default:
			vd.SetAlert(err)
		}
		c.IndexView.Render(w,r,vd)
		return
	}

	err = c.signIn(w, user)
	if err != nil{
		vd.SetAlert(err)
		c.IndexView.Render(w,r,vd)
		return
	}
	http.Redirect(w, r,"/profile", http.StatusFound)
}

func (c Users) Store(w http.ResponseWriter, r *http.Request)  {
	var vd views.Data
	var form RegistrationForm
	if err := helpers.ParseForm(r, &form); err != nil{
		log.Print(err)
		vd.SetAlert(err)
		c.CreateView.Render(w,r,vd)
		return
	}
	user := models.User{
		Name:  form.Name,
		Email: form.Email,
		Password: form.Password,
	}
	err := c.service.Store(&user)
	if err != nil{
		log.Println(err)
		vd.SetAlert(err)
		c.CreateView.Render(w,r,vd)
	    return
	}
	err = c.signIn(w, &user)
	if err != nil{
		log.Println(err.Error())
		http.Redirect(w, r,"/login", http.StatusFound)
	}
	http.Redirect(w, r,"/profile", http.StatusFound)
}

func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.Remember == ""{
		token, err := rand.RememberToken()
		if err != nil{
		    return err
		}
		user.Remember = token
		err = u.service.Update(user)
		if err != nil{
		    return err
		}
	}
	cookie := http.Cookie{
		Name: "remember_token",
		Value: user.Remember,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}