package user

import "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"

type UserService interface {
	Login(auth *model.AuthRequest) (*model.AuthResponse, error)
	Register(user *model.UserEntity) error
}