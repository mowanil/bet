package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"
	"github.com/mowahaeser/bet/database"
	"github.com/mowahaeser/bet/domain"
	"github.com/mowahaeser/bet/inputs"
	"golang.org/x/crypto/bcrypt"
)

var (
	Account IAccount = &account{}
)

type account struct{}

type IAccount interface {
	Login(*inputs.Login) (string, error)
	Register(*inputs.Register) (string, error)
}

func (*account) Register(input *inputs.Register) (string, error) { // (token, err)
	var account domain.Account

	result := database.DB.Where("username = ?", input.Username).First(&account)

	if result.RowsAffected != 0 {
		fmt.Print("entered here")
		return "", errors.New("choose a different username")
	}

	result = database.DB.Where("email = ?", input.Email).First(&account)

	if result.RowsAffected != 0 {
		return "", errors.New("choose a different email address")
	}

	copier.Copy(&account, &input)

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		return "", err
	}

	account.Password = password
	result = database.DB.Create(&account)

	if result.RowsAffected != 0 {
		return "", errors.New("something went wrong")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    account.Id.String(),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})

	token, err := claims.SignedString([]byte("my_secret")) // read from env file

	if err != nil {
		return "", err
	}

	return token, nil
}

func (*account) Login(input *inputs.Login) (string, error) { // token, err

	return "", nil
}
