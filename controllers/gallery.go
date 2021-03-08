package controllers

import (
	"fmt"
	"github.com/thiduzz/lenslocked.com/context"
	"github.com/thiduzz/lenslocked.com/models"
	"github.com/thiduzz/lenslocked.com/views"
	"log"
	"net/http"
)

type Galleries struct {
	CreateView *views.View
	service	models.GalleryService
}

type StoreForm struct {
	Title string `schema:"title"`
}

type UpdateForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		CreateView: views.NewView("master", "galleries/create"),
		service: gs,
	}
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
	fmt.Fprint(w, gallery)
}