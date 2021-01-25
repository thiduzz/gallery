package controllers

import (
	"github.com/gorilla/schema"
	"net/http"
)

func parseForm(r *http.Request, dst interface{}) error  {
	if err := r.ParseForm();err != nil{
		return err
	}
	if err := schema.NewDecoder().Decode(dst, r.PostForm); err != nil{
		return err
	}
	return nil
}
