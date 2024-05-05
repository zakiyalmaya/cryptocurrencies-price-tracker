package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/tracker"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/user"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/middleware"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/transport/controller"
)

func Handlers(userSvc user.UserService, trackerSvc tracker.TrackerService, r *gin.Engine) {
	ctrl := controller.NewController(userSvc, trackerSvc)

	r.POST("/user", ctrl.User.Register)
	r.POST("/user/login", ctrl.User.Login)

	r.Use(middleware.AuthMiddleware()).GET("/user/tracker", ctrl.Tracker.GetUserTrackedList)
	r.Use(middleware.AuthMiddleware()).POST("/user/tracker", ctrl.Tracker.AddUserTrackedCoin)
	r.Use(middleware.AuthMiddleware()).DELETE("/user/tracker/:coinId", ctrl.Tracker.DeleteUserTrackedCoin)
	r.GET("/test", ctrl.Tracker.GetList)
}
