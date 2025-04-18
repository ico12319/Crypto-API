package coin

import (
	"context"
	"crptoApi/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	if err = json.NewDecoder(resp.Body).Decode(&coin); err != nil {
		return 0.0, fmt.Errorf("unavailabe crypto token")
	}
	parsedPrice, _ := strconv.ParseFloat(coin.Price, 64)
	return parsedPrice, nil
}

func formatUrl(cryptoId string) string {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sUSDT", cryptoId)
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
