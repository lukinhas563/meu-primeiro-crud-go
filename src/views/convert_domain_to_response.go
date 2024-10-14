package views

import (
	"github.com/lukinhas563/meu-primeiro-crud-go/src/controller/models/responses"
	"github.com/lukinhas563/meu-primeiro-crud-go/src/models"
)

func ConvertDomainToResponse(userDomain models.UserDomainInterface) responses.UserResponse {
	return responses.UserResponse{
		Id:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
