package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/user"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/transport/controller"
)

func main() {
	db := repository.DBConnection()
	defer db.Close()

	repository := repository.NewRespository(db)

	// initiate service
	userService := user.NewUserService(repository)

	// instatiate router
	r := gin.Default()

	// call handlers
	controller.UserHandlers(userService, r)

	r.Run(":8080")
}