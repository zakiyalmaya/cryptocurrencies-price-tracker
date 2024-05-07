package user

import "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=repo.go -destination=UserRepository.go
type UserRepository interface {
	Create(user *model.UserEntity) error
	Get(username string) (*model.UserEntity, error)
}