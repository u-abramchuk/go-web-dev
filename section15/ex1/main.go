package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/abramchuk/go-web-dev/section15/ex1/controllers"
	"gitlab.com/abramchuk/go-web-dev/section15/ex1/models"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getUserStorage())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8887", r)
}

func getUserStorage() map[string]models.User {
	return map[string]models.User{}
}
