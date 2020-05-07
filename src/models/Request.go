package models

import (
    "strconv"
)

type Request struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
	Hits  int
}

func (r *Request) Content() string {
    return strconv.Itoa(r.Int1) + strconv.Itoa(r.Int2) + strconv.Itoa(r.Limit) + r.Str1 + r.Str2
}
