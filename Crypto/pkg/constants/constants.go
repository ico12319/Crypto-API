package constants

type TransactionType string

const (
	Buy  TransactionType = "buy"
	Sell TransactionType = "sell"
)

var ACCEPTED_TOKENS = map[string]struct{}{
	"bitcoin":     {},
	"dogecoin":    {},
	"ethereum":    {},
	"solana":      {},
	"polkadot":    {},
	"litecoin":    {},
	"polygon":     {},
	"cardano":     {},
	"binancecoin": {},
	"uniswap":     {},
	"stellar":     {},
}

const CONTENT_TYPE = "Content-Type"
const JSON = "application/json"

const USD = "usd"
