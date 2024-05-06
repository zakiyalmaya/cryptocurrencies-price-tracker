package user

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
	return &UserController{
		userSvc: userSvc,
	}
}

func (u *UserController) Register(c *gin.Context) {
	defer c.Request.Body.Close()

	user := &model.UserEntity{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	err := u.userSvc.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(nil))
}

func (u *UserController) Login(c *gin.Context) {
	defer c.Request.Body.Close()

	auth := &model.AuthRequest{}
	if err := c.ShouldBindJSON(auth); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	res, err := u.userSvc.Login(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(res))
}
