package models

import (
	rest_error "github.com/lukinhas563/meu-primeiro-crud-go/src/config/err"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/config/logs"
	"go.uber.org/zap"
)

func (ud *userDomain) GetUser(id string) (*userDomain, *rest_error.RestError) {
	logs.Info("Init getUser model", zap.String("journey", "getUser"))

	return nil, nil
}
