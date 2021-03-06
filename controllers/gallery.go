package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thiduzz/lenslocked.com/context"
	"github.com/thiduzz/lenslocked.com/helpers"
	"github.com/thiduzz/lenslocked.com/models"
	"github.com/thiduzz/lenslocked.com/views"
	"log"
	"net/http"
	"strconv"
)

type Galleries struct {
	CreateView *views.View
	ShowView   *views.View
	EditView   *views.View
	IndexView  *views.View
	service    models.GalleryService
	router     *mux.Router
}

type StoreForm struct {
	Title string `schema:"title"`
}

type UpdateForm struct {
	Title string `schema:"title"`
}

func NewGalleries(gs models.GalleryService, r *mux.Router) *Galleries {
	return &Galleries{
		CreateView: views.NewView("master", "galleries/create"),
		EditView: views.NewView("master", "galleries/edit"),
		ShowView: views.NewView("master", "galleries/show"),
		IndexView: views.NewView("master", "galleries/index"),
		service: gs,
		router: r,
	}
}

func (c *Galleries) Index(w http.ResponseWriter, r *http.Request) {
	user := context.User(r.Context())
	galleries, err := c.service.ByUserID(user.ID)
	if err != nil{
		http.Error(w, "Gallery not found", http.StatusInternalServerError)
	}
	vd := views.Data{
		Alert: nil,
		Yield: map[string] interface{}{"owner": user,"galleries": galleries},
	}
	c.IndexView.Render(w, r, vd)
}


func (c *Galleries) Show(w http.ResponseWriter, r *http.Request)  {
	gallery, err := c.galleryByID(w, r)
	if err != nil{
		http.Error(w, "Gallery not found", http.StatusNotFound)
	}
	c.ShowView.Render(w, r, gallery)
}

func (c *Galleries) Create(w http.ResponseWriter, r *http.Request)  {
	c.CreateView.Render(w, r, map[string] interface{}{"type": "create","gallery": nil})
}

func (c *Galleries) Edit(w http.ResponseWriter, r *http.Request)  {
	gallery, err := c.galleryByID(w, r)
	if err != nil{
		return
	}
	user := context.User(r.Context())
	if gallery.UserID != user.ID{
		http.Error(w, "Gallery not found", http.StatusNotFound)
		return
	}
	c.EditView.Render(w, r, map[string] interface{}{"type": "edit","gallery": gallery})
}

func (c *Galleries) galleryByID(w http.ResponseWriter, r *http.Request) (*models.Gallery, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid gallery ID", http.StatusNotFound)
		return nil, err
	}
	gallery, err := c.service.ByID(uint(id))
	if err != nil {
		switch err {
		case models.ErrNotFound:
			http.Error(w, "Gallery not found", http.StatusNotFound)
		default:
			http.Error(w, "Whoops! Something went wrong!", http.StatusInternalServerError)
		}
		return nil, err
	}
	return gallery, nil
}

func (c *Galleries) Update(w http.ResponseWriter, r *http.Request) {
	gallery, err := c.galleryByID(w, r)
	if err != nil{
		return
	}
	user := context.User(r.Context())
	if gallery.UserID != user.ID{
		http.Error(w, "Gallery not found", http.StatusNotFound)
		return
	}
	var vd views.Data
	vd.Yield = map[string] interface{}{"type": "edit","gallery": gallery}
	var form UpdateForm
	if err := helpers.ParseForm(r, &form); err != nil{
		log.Print(err)
		vd.SetAlert(err)
		c.EditView.Render(w,r,vd)
		return
	}
	gallery.Title = form.Title
	err = c.service.Update(gallery)
	if err != nil{
		log.Println(err)
		vd.SetAlert(err)
		c.EditView.Render(w,r,vd)
		return
	}
	vd.Alert = &views.Alert{
		Color:   views.AlertColorSuccess,
		Title:   "Yeih!",
		Message: "Gallery successfully updated!",
	}
	c.EditView.Render(w,r,vd)
}

func (c *Galleries) Store(w http.ResponseWriter, r *http.Request)  {
	var vd views.Data
	var form StoreForm
	if err := helpers.ParseForm(r, &form); err != nil{
		log.Print(err)
		vd.SetAlert(err)
		c.CreateView.Render(w,r,vd)
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
		c.CreateView.Render(w,r,vd)
		return
	}
	url, err := c.router.Get("gallery.edit").URL("id", fmt.Sprintf("%v",gallery.ID))
	if err != nil{
	    http.Redirect(w,r,"/", http.StatusFound)
	}
	http.Redirect(w, r, url.Path, http.StatusMovedPermanently)
}


func (c *Galleries) Destroy(w http.ResponseWriter, r *http.Request) {
	gallery, err := c.galleryByID(w, r)
	if err != nil{
		return
	}
	user := context.User(r.Context())
	if gallery.UserID != user.ID{
		http.Error(w, "Gallery not found", http.StatusNotFound)
		return
	}
	var vd views.Data
	err = c.service.Destroy(gallery.ID)
	if err != nil{
		log.Println(err)
		vd.SetAlert(err)
		vd.Yield = map[string] interface{}{"type": "edit","gallery": gallery}
		c.EditView.Render(w,r,vd)
		return
	}
	url, _ := c.router.Get("gallery.index").URL()
	http.Redirect(w, r, url.Path, http.StatusFound)
}