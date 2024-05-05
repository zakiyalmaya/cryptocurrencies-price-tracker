package controller

import (
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/tracker"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/user"
	trackerCtrl "github.com/zakiyalmaya/cryptocurrencies-price-tracker/transport/controller/tracker"
	userCtrl "github.com/zakiyalmaya/cryptocurrencies-price-tracker/transport/controller/user"
)

type Controller struct {
	User    *userCtrl.UserController
	Tracker *trackerCtrl.TrackerController
}

func NewController(userSvc user.UserService, trackerSvc tracker.TrackerService) *Controller {
	return &Controller{
		User:    userCtrl.NewUserController(userSvc),
		Tracker: trackerCtrl.NewTrackerController(trackerSvc),
	}
}
