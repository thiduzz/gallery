package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thiduzz/lenslocked.com/controllers"
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
	userService, err := models.NewUserService(psqlInfo)
	if err != nil{
	    panic(err)
	}
	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers(userService)
	router.Handle("/", staticController.Home).Methods(http.MethodGet)
	router.Handle("/contact", staticController.Contact).Methods(http.MethodGet)
	router.HandleFunc("/signup", usersController.Create).Methods(http.MethodGet)
	router.HandleFunc("/signup", usersController.Store).Methods(http.MethodPost)
	router.HandleFunc("/login", usersController.Index).Methods(http.MethodGet)
	router.HandleFunc("/login", usersController.Login).Methods(http.MethodPost)
	router.HandleFunc("/profile", usersController.Show).Methods(http.MethodGet)
	router.NotFoundHandler = http.HandlerFunc(staticController.NotFound)
	http.ListenAndServe(":3000",router)
}