package tracker

import (
	"errors"
	"strconv"

	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client/coincap"
	exchangerate "github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client/exchange_rate"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type trackerSvcImpl struct {
	coinCapSvc      coincap.CoinCapService
	exchangeRateSvc exchangerate.ExchangeRateService
	repos           *repository.Repositories
}

func NewTrackerService(
	coinCapSvc coincap.CoinCapService,
	exchangeRateSvc exchangerate.ExchangeRateService,
	repos *repository.Repositories) TrackerService {
	return &trackerSvcImpl{
		coinCapSvc:      coinCapSvc,
		exchangeRateSvc: exchangeRateSvc,
		repos:           repos,
	}
}

func (t *trackerSvcImpl) GetUserTrackedList(username string) (*model.UserTrackedCoin, error) {
	userCoins, err := t.repos.Tracker.GetUserTrackedCoins(username)
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

	resAssets, err := t.coinCapSvc.GetAssets(&model.AssetRequest{Ids: &coinIds})
	if err != nil {
		return nil, err
	}

	exchangeRate, err := t.exchangeRateSvc.GetLatest("USD", "IDR")
	if err != nil {
		return nil, err
	}

	for _, coin := range userCoins.TrackedCoins {
		for _, asset := range resAssets.Data {
			if asset.ID != coin.CoinID {
				continue
			}

			price, err := strconv.ParseFloat(asset.PriceUsd, 64)
			if err != nil {
				return nil, errors.New("error when parsing price")
			}

			priceIDR := exchangeRate * price
			coin.PriceIDR = &priceIDR
			break
		}
	}

	return userCoins, nil
}

func (t *trackerSvcImpl) AddUserTrackedCoin(req *model.TrackerEntity) error {
	if err := t.repos.Tracker.Create(req); err != nil {
		return err
	}

	return nil
}

func (t *trackerSvcImpl) DeleteUserTrackedCoin(userID int, coinID string) error {
	if err := t.repos.Tracker.Delete(userID, coinID); err != nil {
		return err
	}

	return nil
}

func (t *trackerSvcImpl) GetList() (*[]model.TrackerEntity, error) {
	res, err := t.repos.Tracker.GetList()
	if err != nil {
		return nil, err
	}

	return res, nil
}
