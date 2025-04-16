package models

type Coin struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	MData MarketData `json:"market_data"`
}

type MarketData struct {
	CurrentPrice map[string]float64 `json:"current_price"`
}
