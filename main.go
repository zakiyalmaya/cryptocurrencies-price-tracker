package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/tracker"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/user"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client/coincap"
	exchangerate "github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client/exchange_rate"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/transport"
)

func main() {
	// instatiate repository
	db := repository.DBConnection()
	redcl := repository.RedisClient()
	defer db.Close()

	repository := repository.NewRespository(db, redcl)

	// instatiate client
	clientCoinCap := client.NewAPIClient("https://api.coincap.io")
	clientExchangeRate := client.NewAPIClient("https://api.apilayer.com/exchangerates_data")

	// instatiate service
	userService := user.NewUserService(repository)
	coinCapService := coincap.NewCoinCapService(clientCoinCap)
	exchangeRateService := exchangerate.NewExchangeRateService(clientExchangeRate)
	trackerService := tracker.NewTrackerService(coinCapService, exchangeRateService, repository)

	// instatiate router
	r := gin.Default()

	// call handlers
	transport.Handlers(userService, trackerService, redcl, r)

	r.Run(":8080")
}
