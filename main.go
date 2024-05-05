package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/tracker"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/user"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client/coincap"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/transport"
)

func main() {
	db := repository.DBConnection()
	defer db.Close()

	repository := repository.NewRespository(db)
	client := client.NewAPIClient("https://api.coincap.io")

	// instatiate service
	userService := user.NewUserService(repository)
	coinCapService := coincap.NewCoinCapService(client)
	trackerService := tracker.NewTrackerService(coinCapService, repository)

	// instatiate router
	r := gin.Default()

	// call handlers
	transport.Handlers(userService, trackerService, r)

	r.Run(":8080")
}
