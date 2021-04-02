package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thiduzz/lenslocked.com/models"
	"github.com/thiduzz/lenslocked.com/rand"
	"github.com/thiduzz/lenslocked.com/storage"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
	galleryId, err := c.getGalleryID(r)
	if err != nil {
		c.NewJsonError(w, err)
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		c.NewJsonError(w, err)
		return
	}
	defer file.Close()
	originalName := header.Filename
	s3Location := c.getS3Location(originalName, galleryId)
	url, err := aws.S3PutObject(file, s3Location)
	if err != nil{
		c.NewJsonError(w, err)
		return
	}
	photoId, err := c.service.Store(&models.Photo{
		Title:     r.FormValue("subtitle"),
		Path:      s3Location,
		GalleryID: uint(galleryId),
		Filename: originalName,
	})

	if err != nil{
		c.NewJsonError(w, err)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"id": strconv.Itoa(int(photoId)), "url": url})
}


func (c *Photos) Delete(w http.ResponseWriter, r *http.Request) {
	panic("todo")
}

func (c *Photos) getGalleryID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return 0, errors.New("Invalid gallery ID")
	}
	return id, nil
}

func (c *Photos) getS3Location(originalName string, galleryId int) string  {
	randomString, _ := rand.String(rand.ImageTokenBytes);
	extension := strings.Split(originalName,".")[1:][0]
	location := fmt.Sprintf("%d/%s.%s", galleryId, url.PathEscape(randomString), extension)
	return location
}

func (c *Photos) NewJsonError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}