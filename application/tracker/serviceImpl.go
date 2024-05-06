package tracker

import (
	"database/sql"
	"fmt"
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
		return nil, fmt.Errorf("error when get user tracked coins to database")
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
		return nil, fmt.Errorf("error when get assets to coincap API client")
	}

	exchangeRate, err := t.exchangeRateSvc.GetLatest("USD", "IDR")
	if err != nil {
		return nil, fmt.Errorf("error when get latest exchange rate to client")
	}

	for _, coin := range userCoins.TrackedCoins {
		for _, asset := range resAssets.Data {
			if asset.ID != coin.CoinID {
				continue
			}

			price, err := strconv.ParseFloat(asset.PriceUsd, 64)
			if err != nil {
				return nil, fmt.Errorf("error when parsing price")
			}

			priceIDR := exchangeRate * price
			coin.PriceIDR = &priceIDR
			break
		}
	}

	return userCoins, nil
}

func (t *trackerSvcImpl) AddUserTrackedCoin(req *model.AddUserTrackedCoinRequest) error {
	asset, err := t.coinCapSvc.GetAsset(req.CoinID)
	if err != nil {
		return fmt.Errorf("error when get asset to coincap API client")
	}

	if asset.Data.ID == "" {
		return fmt.Errorf("error the coin is not in the asset list")
	}

	coinTracked, err := t.repos.Tracker.Get(req.UserID, req.CoinID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error when get tracked coin to database")
	}

	if coinTracked != nil {
		return fmt.Errorf("the coin is already tracked")
	}

	createReq := &model.TrackerEntity{
		UserID: req.UserID,
		Coin: model.Coin{
			CoinID:     asset.Data.ID,
			CoinSymbol: asset.Data.Symbol,
			CoinName:   asset.Data.Name,
		},
	}
	if err := t.repos.Tracker.Create(createReq); err != nil {
		return fmt.Errorf("error when add user tracked coin to database")
	}

	return nil
}

func (t *trackerSvcImpl) DeleteUserTrackedCoin(userID int, coinID string) error {
	if err := t.repos.Tracker.Delete(userID, coinID); err != nil {
		return fmt.Errorf("error when remove user tracked coin from database")
	}

	return nil
}

func (t *trackerSvcImpl) GetAssetList(req *model.AssetRequest) (*model.AssetsResponse, error) {
	response, err := t.coinCapSvc.GetAssets(req)
	if err != nil {
		return nil, fmt.Errorf("error when get assets to coincap API client")
	}

	return response, nil
}