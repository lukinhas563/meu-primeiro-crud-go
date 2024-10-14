package service

import (
	rest_error "github.com/lukinhas563/meu-primeiro-crud-go/src/config/err"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/models"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(models.UserDomainInterface) *rest_error.RestError
	UpdateUser(string, models.UserDomainInterface) *rest_error.RestError
	GetUser(string) (*models.UserDomainInterface, *rest_error.RestError)
	DeleteUser(string) *rest_error.RestError
}
