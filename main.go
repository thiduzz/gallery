package main

import (
	"github.com/gorilla/mux"
	"github.com/thiduzz/lenslocked.com/controllers"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers()
	router.Handle("/", staticController.Home).Methods(http.MethodGet)
	router.Handle("/contact", staticController.Contact).Methods(http.MethodGet)
	router.HandleFunc("/signup", usersController.Create).Methods(http.MethodGet)
	router.HandleFunc("/signup", usersController.Store).Methods(http.MethodPost)
	router.HandleFunc("/login", usersController.Index).Methods(http.MethodGet)
	router.NotFoundHandler = http.HandlerFunc(staticController.NotFound)
	http.ListenAndServe(":3000",router)
}