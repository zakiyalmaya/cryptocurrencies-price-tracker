package tracker

import "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=repo.go -destination=TrackerRepository.go
type TrackerRepository interface {
	GetByUsername(username string) (*model.UserTrackedCoin, error)
	Create(req *model.TrackerEntity) error
	Delete(userID int, coinID string) error
	GetByUserIDAndCoinID(userID int, coinID string) (*model.TrackerEntity, error)
}
