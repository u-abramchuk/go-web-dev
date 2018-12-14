package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/abramchuk/go-web-dev/section15/ex2/models"

	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
)

type UserController struct {
	userStore *models.UserStore
}

func NewUserController(s *models.UserStore) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Fetch user
	u, ok := uc.userStore.Find(id)
	if !ok {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	id, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(400)
		return
	}

	u.Id = id.String()

	// store the user in mongodb
	uc.userStore.Add(u)

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	user, ok := uc.userStore.Find(id)
	if !ok {
		w.WriteHeader(404)
		return
	}

	uc.userStore.Remove(user)
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}
