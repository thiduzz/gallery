package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleFunc(rw http.ResponseWriter, r *http.Request)  {
	rw.Header().Set("Content-Type", "text/html")
	switch r.URL.Path {
	case "/":
		fmt.Fprint(rw, "<h1>Main page</h1>")
		break
	case "/contact":
		fmt.Fprint(rw, "To get in touch, send an email to <a href=\"mailto:support@lenslocked.com\">Support Lenslocked.com</a>")
		break
	default:
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprint(rw, "<h1>404</h1>")

	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000",router)
}
