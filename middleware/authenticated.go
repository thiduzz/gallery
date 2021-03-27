package middleware

import (
	"github.com/thiduzz/lenslocked.com/context"
	"net/http"
)

type Authenticated struct {
	SetUser
}

func (mw *Authenticated) Handle(next http.Handler) http.HandlerFunc {
	return mw.HandleFn(next.ServeHTTP)
}

func (mw *Authenticated) HandleFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if context.User(r.Context()) == nil {
			http.Redirect(w,r,"/login",http.StatusFound)
			return
		}
		next(w,r)
	})
}