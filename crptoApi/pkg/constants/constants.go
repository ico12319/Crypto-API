package constants

type TransactionType string

const (
	Buy  TransactionType = "buy"
	Sell TransactionType = "sell"
)

const USD = "usd"

const CONTENT_TYPE = "Content-Type"
const JSON = "application/json"

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