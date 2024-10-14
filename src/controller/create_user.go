package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	rest_error "github.com/lukinhas563/meu-primeiro-crud-go/src/config/err"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/controller/models/requests"
)

func CreateUser(c *gin.Context) {
	var userRequest requests.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errResponse := rest_error.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))

		c.JSON(errResponse.Code, errResponse)
		return
	}

	fmt.Println(userRequest)
}
