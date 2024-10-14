package models

import (
	rest_error "github.com/lukinhas563/meu-primeiro-crud-go/src/config/err"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/logs"
	"go.uber.org/zap"
)

func (ud *userDomain) DeleteUser(id string) *rest_error.RestError {
	logs.Info("Init deleteUser model", zap.String("journey", "deleteUser"))

	return nil
}
