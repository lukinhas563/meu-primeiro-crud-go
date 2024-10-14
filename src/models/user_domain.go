package models

import (
	"crypto/md5"
	"encoding/hex"

	rest_error "github.com/lukinhas563/meu-primeiro-crud-go/src/config/err"
)

func NewUserDomain(
	emain, password, name string,
	age int8,
) *userDomain {
	return &userDomain{
		Email:    emain,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type userDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_error.RestError
	UpdateUser(string) *rest_error.RestError
	GetUset(string) (*userDomain, *rest_error.RestError)
	DeleteUser(string) *rest_error.RestError
}
