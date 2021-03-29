package controllers

import (
	"github.com/thiduzz/lenslocked.com/models"
	"net/http"
)

type Photos struct {
	service    models.PhotoService
}

type StorePhotoForm struct {
	Name string `schema:"name"`
}

func NewPhotos(gs models.PhotoService) *Photos {
	return &Photos{
		service: gs,
	}
}

func (c *Photos) Store(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

