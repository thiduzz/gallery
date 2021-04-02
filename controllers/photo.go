package controllers

import (
	"github.com/thiduzz/lenslocked.com/models"
	"github.com/thiduzz/lenslocked.com/storage"
	"log"
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
	aws := storage.NewAWSConnection()
	filename, err := aws.S3PutObject(r.Body,"test")
	if err != nil{
		log.Println(err)
		return
	}
	panic(filename)
}


func (c *Photos) Delete(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}
