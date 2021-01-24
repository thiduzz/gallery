package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request)   {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Main page</h1>")
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
	w.WriteHeader(404)
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
