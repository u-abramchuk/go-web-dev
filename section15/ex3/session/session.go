package session

import (
	"fmt"
	"time"

	"gitlab.com/abramchuk/go-web-dev/section15/ex3/models"
)

const SessionLength int = 30

type session struct {
	un           string
	lastActivity time.Time
}

type SessionStorage struct {
	users    map[string]models.User
	sessions map[string]session
	cleaned  time.Time
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{
		users:    map[string]models.User{},
		sessions: map[string]session{},
		cleaned:  time.Now(),
	}
}

func (s *SessionStorage) GetUserFromSession(id string) (models.User, bool) {
	if session, ok := s.GetSession(id); ok {
		user, ok := s.users[session.un]

		return user, ok
	}

	return models.User{}, false
}

func (s *SessionStorage) GetUser(id string) (models.User, bool) {
	user, ok := s.users[id]

	return user, ok
}

func (s *SessionStorage) AddUser(user models.User) {
	s.users[user.UserName] = user
}

func (s *SessionStorage) GetSession(id string) (session, bool) {
	session, ok := s.sessions[id]
	if ok {
		session.lastActivity = time.Now()
	}

	return session, ok
}
func (s *SessionStorage) StartSession(sessionId, username string) {
	s.sessions[sessionId] = session{username, time.Now()}
}

func (s *SessionStorage) ClearSession(id string) {
	delete(s.sessions, id)
}

func (s *SessionStorage) CleanSessions() {
	if time.Now().Sub(s.cleaned) > (time.Second * 30) {
		fmt.Println("BEFORE CLEAN") // for demonstration purposes
		s.ShowSessions()            // for demonstration purposes
		for k, v := range s.sessions {
			if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
				delete(s.sessions, k)
			}
		}
		s.cleaned = time.Now()
		fmt.Println("AFTER CLEAN") // for demonstration purposes
		s.ShowSessions()           // for demonstration purposes
	}
}

// var DbUsers = map[string]models.User{} // user ID, user
// var DbSessions = map[string]Session{}  // session ID, session
// var DbSessionsCleaned time.Time

// for demonstration purposes
func (s *SessionStorage) ShowSessions() {
	fmt.Println("********")
	for k, v := range s.sessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
