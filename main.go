package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thiduzz/lenslocked.com/controllers"
	"github.com/thiduzz/lenslocked.com/helpers"
	"github.com/thiduzz/lenslocked.com/middleware"
	"github.com/thiduzz/lenslocked.com/models"
	"net/http"
	"os"
)


func main() {
	router := mux.NewRouter()
	helpers.LoadEnv()
	psqlInfo := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"),  os.Getenv("POSTGRES_PORT"),  os.Getenv("POSTGRES_PASSWORD"),  os.Getenv("POSTGRES_USER"),  os.Getenv("POSTGRES_DB"))
	services, err := models.NewServices(psqlInfo)
	if err != nil {
	    panic(err)
	}
	err = services.AutoMigrate()
	if err != nil {
		panic(err)
	}

	setUserMiddleware := middleware.SetUser{UserService:services.User}
	authenticatedMiddleware := middleware.Authenticated{SetUser: setUserMiddleware}
	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers(services.User)
	galleriesController := controllers.NewGalleries(services.Gallery, router)
	photosController := controllers.NewPhotos(services.Photo)
	router.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))
	router.Handle("/", staticController.Home).Methods(http.MethodGet)
	router.Handle("/contact", staticController.Contact).Methods(http.MethodGet)
	router.HandleFunc("/signup", usersController.Create).Methods(http.MethodGet)
	router.HandleFunc("/signup", usersController.Store).Methods(http.MethodPost)
	router.HandleFunc("/login", usersController.Index).Methods(http.MethodGet)
	router.HandleFunc("/login", usersController.Login).Methods(http.MethodPost)
	router.HandleFunc("/profile", usersController.Show).Methods(http.MethodGet)

	//Galleries
	router.HandleFunc("/galleries", authenticatedMiddleware.HandleFn(galleriesController.Index)).Methods(http.MethodGet).Name("gallery.index")
	router.HandleFunc("/galleries", authenticatedMiddleware.HandleFn(galleriesController.Store)).Methods(http.MethodPost).Name("gallery.store")
	router.HandleFunc("/galleries/create", authenticatedMiddleware.HandleFn(galleriesController.Create)).Methods(http.MethodGet).Name("gallery.create")
	router.HandleFunc("/galleries/{id:[0-9]+}", galleriesController.Show).Methods(http.MethodGet).Name("gallery.show")
	router.HandleFunc("/galleries/{id:[0-9]+}/edit", authenticatedMiddleware.HandleFn(galleriesController.Edit)).Methods(http.MethodGet).Name("gallery.edit")
	router.HandleFunc("/galleries/{id:[0-9]+}/update", authenticatedMiddleware.HandleFn(galleriesController.Update)).Methods(http.MethodPost).Name("gallery.update")
	router.HandleFunc("/galleries/{id:[0-9]+}/destroy", authenticatedMiddleware.HandleFn(galleriesController.Destroy)).Methods(http.MethodPost).Name("gallery.destroy")
	router.HandleFunc("/galleries/{id:[0-9]+}/photos",  authenticatedMiddleware.HandleFn(photosController.Store)).Methods(http.MethodPost).Name("photo.store")
	router.NotFoundHandler = http.HandlerFunc(staticController.NotFound)
	http.ListenAndServe(":3000", setUserMiddleware.Handle(router))
}