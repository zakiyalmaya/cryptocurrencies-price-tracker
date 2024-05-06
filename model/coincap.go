package model

type AssetRequest struct {
	Search *string `json:"search,omitempty"`
	Ids    *string `json:"ids,omitempty"`
	Offset *string `json:"offset,omitempty"`
	Limit  *string `json:"limit,omitempty"`
}

type Response struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
}

type AssetsResponse struct {
	Data []Response `json:"data"`
}

type AssetResponse struct {
	Data Response `json:"data"`
}
