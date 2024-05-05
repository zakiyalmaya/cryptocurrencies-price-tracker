package tracker

import (
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type TrackerService interface {
	GetUserTrackedList(username string) (*model.UserTrackedCoin, error)
	AddUserTrackedCoin(req *model.TrackerEntity) error
	DeleteUserTrackedCoin(userID int, coinID string) error
	GetList() (*[]model.TrackerEntity, error)
}
