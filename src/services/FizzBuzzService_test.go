package services

import (
	"github.com/stretchr/testify/assert"
	"models"
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

func TestGetFizzBuzz(t *testing.T) {

    fizzBuzzService := FizzBuzzService{}

	expectedResult := []string{"1","fizz","buzz","fizz","5","fizzbuzz","7","fizz","buzz","fizz"}

	actualResult, _ := fizzBuzzService.GetFizzBuzz(request)

	assert.Equal(t, expectedResult, actualResult)
}

func TestGetHighestRequest(t *testing.T) {

    fizzBuzzService := FizzBuzzService{}

	actualResult, _ := fizzBuzzService.GetHighestRequest()

	assert.Equal(t, request, actualResult)
}