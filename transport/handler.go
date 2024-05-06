package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/tracker"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/user"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/middleware"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/transport/controller"
)

func Handlers(userSvc user.UserService, trackerSvc tracker.TrackerService, 
	redcl *redis.Client, r *gin.Engine) {
	ctrl := controller.NewController(userSvc, trackerSvc)

	r.POST("/user", ctrl.User.Register)
	r.POST("/user/login", ctrl.User.Login)
	r.Use(middleware.AuthMiddleware(redcl)).POST("/user/logout", ctrl.User.Logout)

	r.Use(middleware.AuthMiddleware(redcl)).GET("/assets", ctrl.Tracker.GetAssetList)
	
	r.Use(middleware.AuthMiddleware(redcl)).GET("/user/tracker", ctrl.Tracker.GetUserTrackedList)
	r.Use(middleware.AuthMiddleware(redcl)).POST("/user/tracker", ctrl.Tracker.AddUserTrackedCoin)
	r.Use(middleware.AuthMiddleware(redcl)).DELETE("/user/tracker/:coinId", ctrl.Tracker.DeleteUserTrackedCoin)
}
