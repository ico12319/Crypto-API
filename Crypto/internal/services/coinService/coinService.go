package coinService

import (
	"CryptoToken/pkg/constants"
	"CryptoToken/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type CoinService interface {
	GetCrypto(cryptoId string) (models.Coin, error)
	GetCryptoPrice(cryptoId string) (float64, error)
}

type HttpCoinService struct {
	client *http.Client
}

func NewHttpCoinService(client *http.Client) *HttpCoinService {
	return &HttpCoinService{client: client}
}

func (h *HttpCoinService) GetCrypto(cryptoId string) (models.Coin, error) {
	url := formatUrl(cryptoId)

	resp, err := getResponse(url, h.client)
	if err != nil {
		return models.Coin{}, err
	}

	var coin models.Coin
	if err := json.NewDecoder(resp.Body).Decode(&coin); err != nil {
		return models.Coin{}, err
	}
	return coin, nil
}

func (h *HttpCoinService) GetCryptoPrice(cryptoId string) (float64, error) {
	url := formatUrl(cryptoId)

	resp, err := getResponse(url, h.client)
	if err != nil {
		return 0.0, err
	}

	var coin models.Coin
	if err := json.NewDecoder(resp.Body).Decode(&coin); err != nil {
		return 0.0, err
	}
	price, ok := coin.MarketData.CurrentPrice[constants.USD]
	if !ok {
		return 0.0, fmt.Errorf("no data for this token")
	}
	return price, nil
}

func formatUrl(cryptoId string) string {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s?localization=false", cryptoId)
	return url
}

func getResponse(url string, client *http.Client) (*http.Response, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
