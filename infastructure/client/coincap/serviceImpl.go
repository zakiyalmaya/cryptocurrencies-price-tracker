package coincap

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type coinCapServiceImpl struct {
	CoinCap *client.APIClient
}

func NewCoinCapService(client *client.APIClient) CoinCapService {
	return &coinCapServiceImpl{
		CoinCap: client,
	}
}

func (c *coinCapServiceImpl) GetAssets(req *model.AssetRequest) (*model.AssetsResponse, error) {
	endpoint := "/v2/assets/"
	url := c.CoinCap.BaseURL + endpoint

	requestBodyJSON, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", url, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return nil, err
	}

	resp, err := c.CoinCap.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &model.AssetsResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}
