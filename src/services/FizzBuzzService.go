package services

import (
    "errors"
	"strconv"
)

type FizzBuzzService struct{}

func (service *FizzBuzzService) GetFizzBuzz(int1 int, int2 int, limit int, str1 string, str2 string) ([]string, error) {

    result := make([]string, 0)

    if limit < 1 {
        return result, errors.New("can't work with limit < 1")
    }

    for i := 1; i <= limit; i++ {
	    result = append(result, ExtractValue(i, int1, int2, str1, str2))
    }

	return result, nil
}

func ExtractValue(i int, int1 int, int2 int, str1 string, str2 string) string {
    result1 := i % int1
    result2 := i % int2

    if (result1 == 0 && result2 == 0) {
        return str1 + str2
    } else if (result1 == 0) {
       return str1
    } else if (result2 == 0) {
        return str2
    }

    return strconv.Itoa(i)
}