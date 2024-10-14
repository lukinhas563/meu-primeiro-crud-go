package service

import (
	rest_error "github.com/lukinhas563/meu-primeiro-crud-go/src/config/err"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/logs"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/models"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain models.UserDomainInterface) *rest_error.RestError {
	logs.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	return nil
}
