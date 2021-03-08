package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thiduzz/lenslocked.com/controllers"
	"github.com/thiduzz/lenslocked.com/middleware"
	"github.com/thiduzz/lenslocked.com/models"
	"net/http"
)

const (
	host = "localhost"
	port = 5434
	user = "admin"
	password = "123456"
	dbname = "lenslocked"
)

func main() {
	router := mux.NewRouter()
	psqlInfo := fmt.Sprintf("host=%s port=%d password=%s user=%s dbname=%s sslmode=disable", host, port, password, user, dbname)
	services, err := models.NewServices(psqlInfo)
	if err != nil {
	    panic(err)
	}
	err = services.AutoMigrate()
	if err != nil {
		panic(err)
	}

	authenticatedMiddleware := middleware.Authenticated{UserService:services.User}
	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers(services.User)
	galleriesController := controllers.NewGalleries(services.Gallery)
	router.Handle("/", staticController.Home).Methods(http.MethodGet)
	router.Handle("/contact", staticController.Contact).Methods(http.MethodGet)
	router.HandleFunc("/signup", usersController.Create).Methods(http.MethodGet)
	router.HandleFunc("/signup", usersController.Store).Methods(http.MethodPost)
	router.HandleFunc("/login", usersController.Index).Methods(http.MethodGet)
	router.HandleFunc("/login", usersController.Login).Methods(http.MethodPost)
	router.HandleFunc("/profile", usersController.Show).Methods(http.MethodGet)

	//Galleries
	router.HandleFunc("/galleries/create", authenticatedMiddleware.HandleFn(galleriesController.Create)).Methods(http.MethodGet)
	router.HandleFunc("/galleries", authenticatedMiddleware.HandleFn(galleriesController.Store)).Methods(http.MethodPost)

	router.NotFoundHandler = http.HandlerFunc(staticController.NotFound)
	http.ListenAndServe(":3000",router)
}