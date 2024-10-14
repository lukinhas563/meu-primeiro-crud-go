package controller

import (
	"github.com/gin-gonic/gin"
	rest_error "github.com/lukinhas563/meu-primeiro-crud-go/src/config/err"
)

func (uc *userControllerInterface) GetUserById(c *gin.Context) {
	err := rest_error.NewBadRequestError("Você chamou essa rota errada")

	c.JSON(err.Code, err)
}

func (uc *userControllerInterface) GetUserByEmail(c *gin.Context) {

}
