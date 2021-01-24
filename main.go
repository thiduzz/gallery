package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thiduzz/lenslocked.com/controllers"
	"github.com/thiduzz/lenslocked.com/views"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request)   {
	must(views.NewView("master", "views/home.gohtml").Render(w,nil))
}

func contact(w http.ResponseWriter, r *http.Request)   {
	must(views.NewView("master", "views/contact.gohtml").Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404</h1>")
}

func main() {
	router := mux.NewRouter()
	usersController := controllers.NewUsers()
	router.HandleFunc("/", home).Methods(http.MethodGet)
	router.HandleFunc("/contact", contact).Methods(http.MethodGet)
	router.HandleFunc("/signup", usersController.New).Methods(http.MethodGet)
	router.HandleFunc("/signup", usersController.Store).Methods(http.MethodPost)
	router.HandleFunc("/login", usersController.Index).Methods(http.MethodGet)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000",router)
}

func must(err error)  {
	if err != nil{
	    panic(err)
	}
}