package main

import (
	"sync"
	"controllers"
	"services"
)

type IServiceContainer interface {
	InjectFizzBuzzController() controllers.FizzBuzzController
}

type kernel struct{}

func (k *kernel) InjectFizzBuzzController() controllers.FizzBuzzController {

	fizzBuzzService := &services.FizzBuzzService{}
	fizzBuzzController := controllers.FizzBuzzController{fizzBuzzService}

	return fizzBuzzController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}