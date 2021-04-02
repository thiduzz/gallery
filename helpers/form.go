package helpers

import (
	"github.com/gorilla/schema"
	"net/http"
)

func ParseForm(r *http.Request, dst interface{}) error  {
	if err := r.ParseForm();err != nil{
		return err
	}
	if err := schema.NewDecoder().Decode(dst, r.PostForm); err != nil{
		return err
	}
	return nil
}
