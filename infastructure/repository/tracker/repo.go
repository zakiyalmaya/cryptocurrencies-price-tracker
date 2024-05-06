package tracker

import "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"

type TrackerRepository interface {
	GetUserTrackedCoins(username string) (*model.UserTrackedCoin, error)
	Create(req *model.TrackerEntity) error
	Delete(userID int, coinID string) error
	Get(userID int, coinID string) (*model.TrackerEntity, error)
}
