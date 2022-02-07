package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mowahaeser/bet/inputs"
	"github.com/mowahaeser/bet/services"
)

var (
	Account IAccount = &account{}
)

type account struct{}

type IAccount interface {
	Login(*gin.Context)
	Register(*gin.Context)
	Logout(*gin.Context)
}

func (*account) Login(ctx *gin.Context) {
	fmt.Println("perform login")
}

func (*account) Register(ctx *gin.Context) {
	input := ctx.MustGet("input").(*inputs.Register)

	token, err := services.Account.Register(input)

	fmt.Println(token, err)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie("cookie", token, 60*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user sucessfully registered in",
	})
}

func (*account) Logout(ctx *gin.Context) {
	fmt.Println("perform logout")
}
