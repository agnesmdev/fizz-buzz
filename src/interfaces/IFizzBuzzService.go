package interfaces

import (
    "models"
)

type IFizzBuzzService interface {
	GetFizzBuzz(request models.Request) ([]string, error)
	GetHighestRequest() (models.Request, error)
}