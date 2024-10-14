package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/validation"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/controller/models/requests"
)

func CreateUser(c *gin.Context) {
	var userRequest requests.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errResponse := validation.ValidateUserError(err)

		c.JSON(errResponse.Code, errResponse)
		return
	}

	fmt.Println(userRequest)
}
