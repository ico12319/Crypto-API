package constants

type TransactionType string

const (
	Buy  TransactionType = "buy"
	Sell TransactionType = "sell"
)

const USD = "usd"

const CONTENT_TYPE = "Content-Type"
const JSON = "application/json"

const AUTH_TOKEN = "admin"

var AUTH_ERROR = map[string]string{
	"error": "unauthorised error",
}
