package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request, p httprouter.Params)   {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Main page</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, p httprouter.Params)   {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, send an email to <a href=\"mailto:support@lenslocked.com\">Support Lenslocked.com</a>")
}

func faq(w http.ResponseWriter, r *http.Request, p httprouter.Params)   {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ Page</h1>")
}

func notFound(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(404)
	fmt.Fprint(w, "<h1>404</h1>")
}

func main() {
	router := httprouter.New()
	router.GET("/", home)
	router.GET("/contact", contact)
	router.GET("/faq", faq)
	router.NotFound = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000",router)
}
