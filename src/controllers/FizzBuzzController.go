package controllers

import (
    "encoding/json"
    "errors"
	"interfaces"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type FizzBuzzController struct {
	interfaces.IFizzBuzzService
}

func (controller *FizzBuzzController) ApplyFizzBuzz(writer http.ResponseWriter, req *http.Request) {
    query := req.URL.Query()
    encoder := json.NewEncoder(writer)

    int1Param, errInt1 := ExtractQueryParam(query, "int1")
    int2Param, errInt2 := ExtractQueryParam(query, "int2")
    limitParam, errLimit := ExtractQueryParam(query, "limit")
	str1, _ := ExtractQueryParam(query, "str1")
	str2, _ := ExtractQueryParam(query, "str2")

	if errInt1 != nil {
        ManageError(writer, errInt1, encoder)
        return
	}

	if errInt2 != nil {
        ManageError(writer, errInt2, encoder)
        return
	}

	if errLimit != nil {
        ManageError(writer, errLimit, encoder)
        return
	}

	int1, err1 := strconv.Atoi(int1Param)
	int2, err2 := strconv.Atoi(int2Param)
	limit, errLimit := strconv.Atoi(limitParam)

	if err1 != nil {
        ManageError(writer, err1, encoder)
        return
	}

	if err2 != nil {
        ManageError(writer, err2, encoder)
        return
	}

	if errLimit != nil {
        ManageError(writer, errLimit, encoder)
        return
	}

	fizzBuzz, err := controller.GetFizzBuzz(int1, int2, limit, str1, str2)

	if err != nil {
        ManageError(writer, err, encoder)
        return
	}

    writer.WriteHeader(200)
	encoder.Encode(fizzBuzz)
}

func ManageError(writer http.ResponseWriter, e error, encoder *json.Encoder) {
    writer.WriteHeader(400)
    encoder.Encode(e.Error())
}

func ExtractQueryParam(query url.Values, name string) (string, error) {
    param, present := query[name]

    if !present || len(param) == 0 {
        return "", errors.New("missing parameter ${param}")
    }

    return strings.Join(param, ","), nil
}