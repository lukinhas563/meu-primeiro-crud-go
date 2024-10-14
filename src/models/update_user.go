package models

import (
	rest_error "github.com/lukinhas563/meu-primeiro-crud-go/src/config/err"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/logs"
	"go.uber.org/zap"
)

func (ud *userDomain) UpdateUser() *rest_error.RestError {
	logs.Info("Init updateUser model", zap.String("journey", "updateUser"))

	return nil
}
