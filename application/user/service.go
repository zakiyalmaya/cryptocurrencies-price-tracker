package user

import "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=service.go -destination=UserService.go
type UserService interface {
	Login(auth *model.AuthRequest) (*model.AuthResponse, error)
	Register(user *model.UserEntity) error
	Logout(username string) error
}