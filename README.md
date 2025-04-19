# Crypto-API
A web API for managing a crypto portfolio: fetching real-time prices, controlling account balance, holdings, and transactions.
## 📦 Tech Stack
  Go
  
  SQLite (github.com/mattn/go-sqlite3)
  
  Gorilla Mux (HTTP routing)

## 🚀 Installation
  git clone https://github.com/ico12319/crypto-api.git
  
  cd crypto-api
  
  go mod download
  
  go build -o crypto-server cmd/server.go
  ./crypto-server  # listens on :5050

## 🌐 Endpoints
- **Account**
  
  GET  /account/               # get current balance
  PUT  /account/{quantity}     # update balance

- **Holding**

  GET  /holding/               # return all of your current holdings
  GET  /holding/{crypto_id}    # return specific holding you already own

- **Transaction**

  GET    /transaction/          # return all transactions
  GET    /transaction/{id}      # return transaction by ID
  POST   /transaction/          # create transaction (buy/sell)
  
  Example:
  { "type":"buy", "crypto":"BTC", "quantity":1.5 }

## ⚙️ Configuration
  Port: 5050 (in cmd/server.go)
  
  Cache TTL: 2m (internal/cache/priceCache.go)
  
  Auth header: Authorization: admin (middlewares/validationMiddleware.go), if not provided you won't be authorized to make any requests

## 🏗 Design Patterns
  Decorator – Middlewares implemented as wrappers (Validation, Logging, Content-Type).

  Singleton – In-memory cache (priceCache.go) ensures a single instance.

  Repository – Abstraction of persistence layer (internal/*DB.go).

  Adapter – Converters map between DB entities and API models.

  Dependency Injection – All layers receive dependencies via constructors.










