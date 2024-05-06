package exchangerate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/client"
)

type exchangeRateSvcImpl struct {
	exchangeRate *client.APIClient
	cache        map[string]cachedExchangeRate
	mu           sync.RWMutex
}

type cachedExchangeRate struct {
	Rate      float64
	ExpiredAt time.Time
}

func NewExchangeRateService(client *client.APIClient) ExchangeRateService {
	return &exchangeRateSvcImpl{
		exchangeRate: client,
		cache:        make(map[string]cachedExchangeRate),
	}
}

func (e *exchangeRateSvcImpl) GetLatest(base, to string) (float64, error) {
	cachedExchangeRate := e.getLatestFromCache(base, to)
	if cachedExchangeRate != 0 {
		return cachedExchangeRate, nil
	}

	endpoint := "/latest"
	urlValue := e.exchangeRate.BaseURL + endpoint

	queryParam := "?symbols=" + to + "&base=" + base
	urlValue += queryParam

	request, err := http.NewRequest("GET", urlValue, nil)
	if err != nil {
		return 0, err
	}
	request.Header.Set("apikey", "CgWMWUg82yzKxdiuZ6UYFJWFrUKjzllZ")

	resp, err := e.exchangeRate.Client.Do(request)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	if isSuccess, ok := data["success"]; ok && !isSuccess.(bool) {
		return 0, fmt.Errorf("failed to get exchange rates")
	}

	rates, ok := data["rates"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("failed to parse exchange rates")
	}

	exchangeRate, ok := rates[to].(float64)
	if !ok {
		return 0, fmt.Errorf("exchange rate for %s not found", to)
	}

	e.setLatestToCache(exchangeRate, base, to)
	return exchangeRate, nil
}

func (e *exchangeRateSvcImpl) getLatestFromCache(base, to string) float64 {
	e.mu.Lock()
	cachedRate, ok := e.cache[cacheKey(base, to)]
	e.mu.Unlock()
	if ok && !cachedRate.ExpiredAt.Before(time.Now()) {
		return cachedRate.Rate
	}

	return 0
}

func (e *exchangeRateSvcImpl) setLatestToCache(exchangeRate float64, base, to string) {
	e.mu.Lock()
	e.cache[cacheKey(base, to)] = cachedExchangeRate{
		Rate:      exchangeRate,
		ExpiredAt: time.Now().Add(time.Hour), // Set TTL to 1 hour
	}
	e.mu.Unlock()
}

func cacheKey(base, to string) string {
	return base + "_" + to
}
