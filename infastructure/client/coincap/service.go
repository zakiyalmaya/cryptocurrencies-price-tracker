package coincap

import "github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=service.go -destination=CoinCapService.go
type CoinCapService interface {
	GetAssets(req *model.AssetRequest) (*model.AssetsResponse, error)
	GetAsset(coinID string) (*model.AssetResponse, error)
}