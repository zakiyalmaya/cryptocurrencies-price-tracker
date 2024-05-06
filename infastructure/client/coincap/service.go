package coincap

import "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"

type CoinCapService interface {
	GetAssets(req *model.AssetRequest) (*model.AssetsResponse, error)
	GetAsset(coinID string) (*model.AssetResponse, error)
}