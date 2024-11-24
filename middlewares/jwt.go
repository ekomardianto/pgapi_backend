package middlewares

import (
	"context"
	"eko/api-pg-bpr/helper"
	"net/http"
	"strings"
)

var ResponseFailed = helper.ResponseFailed

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		if accessToken == "" {
			ResponseFailed(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		tokenString := strings.Split(accessToken, " ")[1]
		if tokenString == "" {
			ResponseFailed(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		user, err := helper.ValidateToken(tokenString)
		// var username = user.Username
		if err != nil {
			ResponseFailed(w, http.StatusUnauthorized, err.Error())
			return
		}
		// fmt.Println(username)
		ctx := context.WithValue(r.Context(), "userInfo", user)
		next.ServeHTTP(w, r.WithContext(ctx))

	})

}
