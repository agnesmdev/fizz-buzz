package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"models"
	"net/http/httptest"
	"services"
	"testing"
)

var request = models.Request{
                Int1: 2,
                Int2: 3,
                Limit: 10,
                Str1: "fizz",
                Str2: "buzz",
                Hits: 1,
            }

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

func TestGetStatistics(t *testing.T) {

	fizzBuzzService := new(services.FizzBuzzService)
	fizzBuzzController := FizzBuzzController{fizzBuzzService}

	req := httptest.NewRequest("GET", "http://localhost:8080/statistics", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/statistics", fizzBuzzController.GetStatistics)
	r.ServeHTTP(w, req)

	actualResult := models.Request{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	assert.Equal(t, request, actualResult)
}