package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/err"
)

func GetUserById(c *gin.Context) {
	err := err.NewBadRequestError("Você chamou essa rota errada")

	c.JSON(err.Code, err)
}

func GetUserByEmail(c *gin.Context) {

}
