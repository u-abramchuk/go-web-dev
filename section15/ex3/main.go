package main

import (
	"net/http"

	"gitlab.com/abramchuk/go-web-dev/section15/ex3/controllers"
	"gitlab.com/abramchuk/go-web-dev/section15/ex3/session"
)

func main() {
	uc := controllers.NewUserController(session.NewSessionStorage())
	http.HandleFunc("/", uc.Index)
	http.HandleFunc("/bar", uc.Bar)
	http.HandleFunc("/signup", uc.Signup)
	http.HandleFunc("/login", uc.Login)
	http.HandleFunc("/logout", uc.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8887", nil)
}
