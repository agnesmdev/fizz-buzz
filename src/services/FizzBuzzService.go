package services

import (
    "errors"
    "models"
    "sort"
	"strconv"
)

type FizzBuzzService struct{}

var requests = make([]models.Request, 0)

func (service *FizzBuzzService) GetFizzBuzz(request models.Request) ([]string, error) {

    AppendRequest(request)
    result := make([]string, 0)

    if request.Limit < 1 {
        return result, errors.New("can't work with limit < 1")
    }

    for i := 1; i <= request.Limit; i++ {
	    result = append(result, ExtractValue(i, request.Int1, request.Int2, request.Str1, request.Str2))
    }

	return result, nil
}

func (service *FizzBuzzService) GetHighestRequest() (models.Request, error) {
    if len(requests) == 0 {
        return models.Request{}, errors.New("no statistics available")
    }

    sort.SliceStable(requests, func(i, j int) bool {
        return requests[i].Hits > requests[j].Hits
    })

    return requests[0], nil
}

func AppendRequest(request models.Request) {
    for i, r := range requests {
        if r.Content() == request.Content() {
            requests[i].Hits = requests[i].Hits + 1
            return
        }
    }

    requests = append(requests, request)
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