package tracker

import (
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client/coincap"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type trackerSvc struct {
	coinCapSvc coincap.CoinCapService
	repos      *repository.Repositories
}

func NewTrackerService(coinCapSvc coincap.CoinCapService, repos *repository.Repositories) TrackerService {
	return &trackerSvc{
		coinCapSvc: coinCapSvc,
		repos:      repos,
	}
}

func (c *trackerSvc) GetUserTrackedList(username string) (*model.UserTrackedCoin, error) {
	userCoins, err := c.repos.Tracker.GetUserTrackedCoins(username)
	if err != nil {
		return nil, err
	}

	var coinIds string
	for i, coin := range userCoins.TrackedCoins {
		if i == len(userCoins.TrackedCoins)-1 {
			coinIds += coin.CoinID
		} else {
			coinIds += coin.CoinID + ","
		}
	}

	if coinIds == "" {
		return userCoins, nil
	}

	resAssets, err := c.coinCapSvc.GetAssets(&model.AssetRequest{Ids: &coinIds})
	if err != nil {
		return nil, err
	}

	for _, coin := range userCoins.TrackedCoins {
		for _, asset := range resAssets.Data {
			if asset.ID == coin.CoinID {
				// todo convert USD to IDR
			}
		}
	}

	return userCoins, nil
}

func (c *trackerSvc) AddUserTrackedCoin(req *model.TrackerEntity) error {
	if err := c.repos.Tracker.Create(req); err != nil {
		return err
	}

	return nil
}

func (c *trackerSvc) DeleteUserTrackedCoin(userID int, coinID string) error {
	if err := c.repos.Tracker.Delete(userID, coinID); err != nil {
		return err
	}
	
	return nil
}

func (c *trackerSvc) GetList() (*[]model.TrackerEntity, error) {
	res, err := c.repos.Tracker.GetList()
	if err != nil {
		return nil, err
	}

	return res, nil
}
