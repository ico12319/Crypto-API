package coin

import (
	"context"
	"crptoApi/pkg/constants"
	"crptoApi/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPCoinService struct {
	client *http.Client
}

func NewHttpCoinService(client *http.Client) *HTTPCoinService {
	return &HTTPCoinService{client: client}
}

func (h *HTTPCoinService) GetCoinPrice(ctx context.Context, cryptoId string) (float64, error) {
	url := formatUrl(cryptoId)
	resp, err := getResponse(ctx, h.client, url)
	if err != nil {
		return 0.0, err
	}
	var coin models.Coin
	if err := json.NewDecoder(resp.Body).Decode(&coin); err != nil {
		return 0.0, err
	}
	price, ok := coin.MData.CurrentPrice[constants.USD]
	if !ok {
		return 0.0, fmt.Errorf("no data for this crypto_token %s", cryptoId)
	}
	return price, nil
}

func formatUrl(cryptoId string) string {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s?localization=false", cryptoId)
	return url
}

func getResponse(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
