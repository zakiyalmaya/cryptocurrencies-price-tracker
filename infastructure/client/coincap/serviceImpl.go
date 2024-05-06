package coincap

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/model"
)

type coinCapSvcImpl struct {
	CoinCap *client.APIClient
}

func NewCoinCapService(client *client.APIClient) CoinCapService {
	return &coinCapSvcImpl{
		CoinCap: client,
	}
}

func (c *coinCapSvcImpl) GetAssets(req *model.AssetRequest) (*model.AssetsResponse, error) {
	endpoint := "/v2/assets"
	urlValue := c.CoinCap.BaseURL + endpoint

	queryParam := buildQueryParamGetAssets(req)
	urlValue += "?" + queryParam

	request, err := http.NewRequest("GET", urlValue, nil)
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

func buildQueryParamGetAssets(req *model.AssetRequest) string {
	params := url.Values{}
	if req.Ids != nil {
		params.Set("ids", *req.Ids)
	}

	if req.Search != nil {
		params.Set("search", *req.Search)
	}

	if req.Limit != nil {
		params.Set("limit", *req.Limit)
	}

	if req.Offset != nil {
		params.Set("offset", *req.Offset)
	}

	return params.Encode()
}
