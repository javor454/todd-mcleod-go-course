package exercises

import (
	"fmt"
	"log"
)

type User struct {
	id   int
	name string
}

type MockUserRepository struct {
	users map[int]User
}

func (mr MockUserRepository) GetUser(id int) (User, error) {
	user, ok := mr.users[id]
	if !ok {
		return User{}, fmt.Errorf("user s id (%v) neexistuje\n", id)
	}
	return user, nil
}

func (mr MockUserRepository) SaveUser(u User) error {
	if _, ok := mr.users[u.id]; ok {
		return fmt.Errorf("user s id (%v) jiz existuje\n", u.id)
	}
	mr.users[u.id] = u
	return nil
}

type UserRepository interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type UserService struct {
	ur UserRepository
}

func (s UserService) GetUser(id int) (User, error) {
	return s.ur.GetUser(id)
}

func (s UserService) SaveUser(u User) error {
	return s.ur.SaveUser(u)
}

func Ex11() {
	mr := MockUserRepository{users: map[int]User{}}
	service := UserService{mr}

	u := User{
		id:   1,
		name: "Karel",
	}
	err := service.SaveUser(u)
	if err != nil {
		log.Fatalf("chyba pri ukladani uzivatele: %v\n", err)
	}

	u2, err := service.GetUser(1)
	if err != nil {
		log.Fatalf("chyba pri stazeni uzivatele: %v\n", err)
	}
	fmt.Println(u2)

	fmt.Println("Exercise 11 END------------")
}
