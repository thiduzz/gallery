package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thiduzz/lenslocked.com/context"
	"github.com/thiduzz/lenslocked.com/models"
	"github.com/thiduzz/lenslocked.com/views"
	"log"
	"net/http"
	"strconv"
)

type Galleries struct {
	CreateView *views.View
	ShowView *views.View
	service	models.GalleryService
	router *mux.Router
}

type StoreForm struct {
	Title string `schema:"title"`
}

type UpdateForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}

func NewGalleries(gs models.GalleryService, r *mux.Router) *Galleries {
	return &Galleries{
		CreateView: views.NewView("master", "galleries/create"),
		ShowView: views.NewView("master", "galleries/show"),
		service: gs,
		router: r,
	}
}

func (c *Galleries) Show(w http.ResponseWriter, r *http.Request)  {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		http.Error(w, "Invalid gallery ID", http.StatusNotFound)
		return
	}
	gallery, err := c.service.ByID(uint(id))
	if err != nil{
		switch err {
		case models.ErrNotFound:
			http.Error(w, "Gallery not found", http.StatusNotFound)
			return
		default:
			http.Error(w, "Whoops! Something went wrong!", http.StatusInternalServerError)
			return

		}
	}
	c.ShowView.Render(w, gallery)
}

func (c *Galleries) Create(w http.ResponseWriter, r *http.Request)  {

	c.CreateView.Render(w, nil)
}

func (c *Galleries) Store(w http.ResponseWriter, r *http.Request)  {
	var vd views.Data
	var form StoreForm
	if err := parseForm(r, &form); err != nil{
		log.Print(err)
		vd.SetAlert(err)
		c.CreateView.Render(w,vd)
		return
	}
	gallery := models.Gallery{
		Title:  form.Title,
		UserID: context.User(r.Context()).ID,
	}
	err := c.service.Store(&gallery)
	if err != nil{
		log.Println(err)
		vd.SetAlert(err)
		c.CreateView.Render(w,vd)
		return
	}
	url, err := c.router.Get("gallery.show").URL("id", fmt.Sprintf("%v",gallery.ID))
	if err != nil{
	    http.Redirect(w,r,"/", http.StatusFound)
	}
	http.Redirect(w, r, url.Path, http.StatusMovedPermanently)
}