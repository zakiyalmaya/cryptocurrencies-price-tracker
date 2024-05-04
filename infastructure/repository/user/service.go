package user

import "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"

type UserRepository interface {
	Create(user *model.UserEntity) error
	Get(username string) (*model.UserEntity, error)
}