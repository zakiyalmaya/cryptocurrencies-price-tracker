package model

type TrackerEntity struct {
	ID     int `json:"id"`
	UserID int `json:"userId"`
	Coin
}

type Coin struct {
	CoinID     string  `json:"coinId"`
	CoinSymbol string  `json:"coinSymbol"`
	CoinName   string  `json:"coinName"`
	PriceIDR   float64 `json:"priceIDR"`
}

type UserTrackedCoin struct {
	Username     string `json:"username"`
	TrackedCoins []Coin
}
