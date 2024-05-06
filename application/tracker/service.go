package tracker

import (
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type TrackerService interface {
	GetUserTrackedList(username string) (*model.UserTrackedCoin, error)
	AddUserTrackedCoin(req *model.AddUserTrackedCoinRequest) error
	DeleteUserTrackedCoin(userID int, coinID string) error
	GetAssetList(req *model.AssetRequest) (*model.AssetsResponse, error)
}
