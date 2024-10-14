package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/models/service"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	GetUserById(c *gin.Context)
	GetUserByEmail(c *gin.Context)

	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
