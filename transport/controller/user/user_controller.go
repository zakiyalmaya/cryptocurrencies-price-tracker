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

	user := &model.UserRequest{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	if err := validateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse(err.Error()))
		return
	}

	err := u.userSvc.Register(&model.UserEntity{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
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

	if auth.Username == "" || auth.Password == "" {
		c.JSON(http.StatusBadRequest, model.HTTPErrorResponse("username/password is required!"))
		return
	}

	res, err := u.userSvc.Login(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(res))
}

func (u *UserController) Logout(c *gin.Context) {
	defer c.Request.Body.Close()

	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.HTTPErrorResponse("Username not found in context"))
		return
	}

	if err := u.userSvc.Logout(username.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, model.HTTPErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.HTTPSuccessResponse(nil))
}
