package tracker

import (
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=service.go -destination=TrackerService.go
type TrackerService interface {
	GetUserTrackedList(username string) (*model.UserTrackedCoin, error)
	AddUserTrackedCoin(req *model.AddUserTrackedCoinRequest) error
	DeleteUserTrackedCoin(userID int, coinID string) error
	GetAssetList(req *model.AssetRequest) (*model.AssetsResponse, error)
}
