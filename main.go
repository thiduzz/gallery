package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Address struct {
	Street string
	HouseNumber int
}

type User struct {
	Address
	Name string
	Age int
	Skills []string
	Dog string
	Location map[string] float64
}

func home(w http.ResponseWriter, r *http.Request)   {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("views/home.gohtml")
	if err != nil{
	    panic(err)
	}
	data := User{
		Name:    "Thiago",
		Age:     30,
		Address: Address{
			Street: "Okerstrasse",
			HouseNumber: 35,
		},
		Skills: []string{
			"PHP",
			"MySQL",
			"Leadership",
		},
		Dog:     "Odin",
		Location: map[string]float64{
			"Latitude": 64.123,
			"Longitude": 123.23333,
		},
	}
	err = t.Execute(w, data)
	if err != nil{
	    panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request)   {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, send an email to <a href=\"mailto:support@lenslocked.com\">Support Lenslocked.com</a>")
}

func faq(w http.ResponseWriter, r *http.Request)   {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ Page</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404</h1>")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000",router)
}
