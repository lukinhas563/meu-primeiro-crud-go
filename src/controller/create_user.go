package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/logs"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/validation"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/controller/models/requests"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/models"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/views"
	"go.uber.org/zap"
)

var (
	userDomainInterface models.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logs.Info("Init create_user controller",
		zap.String("journey", "createUser"))

	var userRequest requests.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logs.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))

		errResponse := validation.ValidateUserError(err)

		c.JSON(errResponse.Code, errResponse)
		return
	}

	domain := models.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logs.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, views.ConvertDomainToResponse(domain))
}
