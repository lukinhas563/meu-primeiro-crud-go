package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/logs"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/validation"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/controller/models/requests"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/controller/models/responses"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logs.Info("Init create_user controller",
		zap.String("journey", "createUser"))

	var userRequest requests.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logs.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))

		errResponse := validation.ValidateUserError(err)

		c.JSON(errResponse.Code, errResponse)
		return
	}

	response := responses.UserResponse{
		Id:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logs.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
