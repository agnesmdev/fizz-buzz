package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFizzBuzz(t *testing.T) {

    fizzBuzzService := FizzBuzzService{}

	expectedResult := []string{"1","fizz","buzz","fizz","5","fizzbuzz","7","fizz","buzz","fizz"}

	actualResult, _ := fizzBuzzService.GetFizzBuzz(2, 3, 10, "fizz", "buzz")

	assert.Equal(t, expectedResult, actualResult)
}