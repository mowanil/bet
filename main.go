package main

import (
	"fmt"

	"github.com/mowahaeser/bet/application"
	"github.com/mowahaeser/bet/database"
	"github.com/mowahaeser/bet/domain"
)

func main() {
	fmt.Println("Hello Mars!")

	database.DB.AutoMigrate(&domain.Account{})

	application.Listen()
}
