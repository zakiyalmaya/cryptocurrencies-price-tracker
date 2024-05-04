package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/application/user"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type UserController struct {
	userSvc user.UserService
}

func NewUserController(userSvc user.UserService) *UserController {
	return &UserController{userSvc}
}

func UserHandlers(userSvc user.UserService, r *gin.Engine) {
	userCtrl := NewUserController(userSvc)

	r.POST("/user", userCtrl.Register)
	r.POST("/user/login", userCtrl.Login)
}

func (userCtrl *UserController) Register(c *gin.Context) {
	defer c.Request.Body.Close()

	user := &model.UserEntity{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := userCtrl.userSvc.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Successfully registered")
}

func (userCtrl *UserController) Login(c *gin.Context) {
	defer c.Request.Body.Close()

	auth := &model.AuthRequest{}
	if err := c.ShouldBindJSON(auth); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := userCtrl.userSvc.Login(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}