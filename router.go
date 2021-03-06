package main

import (
	"sync"
	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	fizzBuzzController := ServiceContainer().InjectFizzBuzzController()

	r := chi.NewRouter()
	r.HandleFunc("/fizz-buzz", fizzBuzzController.ApplyFizzBuzz)
	r.HandleFunc("/statistics", fizzBuzzController.GetStatistics)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}