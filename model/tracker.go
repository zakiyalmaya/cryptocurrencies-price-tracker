package model

import "time"

type TrackerEntity struct {
	ID     int `json:"id"`
	UserID int `json:"userId"`
	Coin
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Coin struct {
	CoinID     string   `json:"coinId"`
	CoinSymbol string   `json:"coinSymbol"`
	CoinName   string   `json:"coinName"`
	PriceIDR   *float64 `json:"priceIDR"`
}

type UserTrackedCoin struct {
	Username     string `json:"username"`
	TrackedCoins []*Coin
}

type AddUserTrackedCoinRequest struct {
	UserID int    `json:"userId"`
	CoinID string `json:"coinId"`
}
