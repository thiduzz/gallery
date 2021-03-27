package middleware

import (
	"github.com/thiduzz/lenslocked.com/context"
	"github.com/thiduzz/lenslocked.com/models"
	"net/http"
)

type SetUser struct {
	models.UserService
}

func (mw *SetUser) Handle(next http.Handler) http.HandlerFunc {
	return mw.HandleFn(next.ServeHTTP)
}

func (mw *SetUser) HandleFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("remember_token")
		ctx := r.Context()
		if err != nil {
			ctx = context.WithUser(ctx, nil)
			r = r.WithContext(ctx)
			next(w,r)
		}
		user, err := mw.UserService.ByRemember(cookie.Value)
		if err != nil {
			next(w,r)
		}
		ctx = context.WithUser(ctx, user)
		r = r.WithContext(ctx)
		next(w,r)
	})
}