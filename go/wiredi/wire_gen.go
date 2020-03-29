// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"fmt"
	"github.com/google/wire"
	"log"
)

// Injectors from main.go:

func InitApp() (*Service, error) {
	mainRepo := NewRepo()
	service, err := NewService(mainRepo)
	if err != nil {
		return nil, err
	}
	return service, nil
}

// main.go:

type User struct {
	ID int64
}

type Repo interface {
	Get(int64) (*User, error)
}

type repo struct{}

func (r *repo) Get(id int64) (*User, error) {
	return &User{ID: id}, nil
}

func NewRepo() Repo {
	return &repo{}
}

var (
	Set = wire.NewSet(
		NewRepo,
	)
)

type Service struct {
	R Repo
}

func NewService(r Repo) (*Service, error) {
	return &Service{R: r}, nil
}

func main() {
	s, err := InitApp()
	if err != nil {
		log.Fatal(err)
		return
	}
	u, err := s.R.Get(1)
	fmt.Printf("%+v,+%v", u, err)
}
