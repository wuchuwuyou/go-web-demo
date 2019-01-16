package controller

import (
	"log"
	"net/http"
)

func middleAuth(next http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter,r *http.Request) {
		username,err := getSessionUser(r)
		log.Println("middle:",username)
		if err != nil {
			log.Println("middle get session err and redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		}else {
			next.ServeHTTP(w, r)
		}
	}
}