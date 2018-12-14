package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type UserStore struct {
	users map[string]User
	path  string
}

func NewUserStore(path string) *UserStore {
	users, err := LoadFromFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	return &UserStore{
		users: users,
		path:  path,
	}
}

func (s *UserStore) Add(u User) {
	s.users[u.Id] = u

	SaveToFile(s.path, s.users)
}

func (s *UserStore) Remove(u User) {
	delete(s.users, u.Id)

	SaveToFile(s.path, s.users)
}

func (s *UserStore) Find(id string) (User, bool) {
	user, ok := s.users[id]

	return user, ok
}

func LoadFromFile(path string) (map[string]User, error) {
	users := map[string]User{}

	if _, err := os.Stat(path); os.IsExist(err) {
		content, err := ioutil.ReadFile(path)
		if err == nil {
			return nil, err
		}

		err = json.Unmarshal(content, &users)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func SaveToFile(path string, users map[string]User) error {
	content, err := json.Marshal(users)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}

	return nil
}
