package controllers

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"gitlab.com/abramchuk/go-web-dev/section15/ex3/models"
	"gitlab.com/abramchuk/go-web-dev/section15/ex3/session"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type UserController struct {
	session *session.SessionStorage
}

func NewUserController(s *session.SessionStorage) *UserController {
	return &UserController{s}
}

func (us *UserController) getUser(w http.ResponseWriter, req *http.Request) models.User {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = session.SessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	u, _ := us.session.GetUserFromSession(c.Value)

	return u
}

func (uс *UserController) alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	_, ok := uс.session.GetUserFromSession(c.Value)
	c.MaxAge = session.SessionLength
	http.SetCookie(w, c)
	return ok
}

func (uс UserController) Index(w http.ResponseWriter, req *http.Request) {
	u := uс.getUser(w, req)
	uс.session.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (uc UserController) Bar(w http.ResponseWriter, req *http.Request) {
	u := uc.getUser(w, req)
	if !uc.alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	uc.session.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func (uc UserController) Signup(w http.ResponseWriter, req *http.Request) {
	if uc.alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := uc.session.GetUser(un); ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = session.SessionLength
		http.SetCookie(w, c)
		uc.session.StartSession(c.Value, un)
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{un, bs, f, l, r}
		uc.session.AddUser(u)
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	uc.session.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func (uc UserController) Login(w http.ResponseWriter, req *http.Request) {
	if uc.alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := uc.session.GetUser(un)
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = session.SessionLength
		http.SetCookie(w, c)
		uc.session.StartSession(c.Value, un)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	uc.session.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func (uc UserController) Logout(w http.ResponseWriter, req *http.Request) {
	if !uc.alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	uc.session.ClearSession(c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions

	go uc.session.CleanSessions()

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
