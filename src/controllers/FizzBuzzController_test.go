package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"services"
	"testing"
)

func TestApplyFizzBuzz(t *testing.T) {

	fizzBuzzService := new(services.FizzBuzzService)
	fizzBuzzController := FizzBuzzController{fizzBuzzService}

	req := httptest.NewRequest("GET", "http://localhost:8080/fizz-buzz?int1=2&int2=3&limit=10&str1=fizz&str2=buzz", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/fizz-buzz", fizzBuzzController.ApplyFizzBuzz)
	r.ServeHTTP(w, req)

	expectedResult := []string{"1","fizz","buzz","fizz","5","fizzbuzz","7","fizz","buzz","fizz"}

	actualResult := make([]string, 0)
	json.NewDecoder(w.Body).Decode(&actualResult)

	assert.Equal(t, expectedResult, actualResult)
}