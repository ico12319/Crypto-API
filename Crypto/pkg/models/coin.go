package models

type Coin struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MarketData struct {
		CurrentPrice map[string]float64 `json:"current_price"`
	} `json:"market_data"`
}
