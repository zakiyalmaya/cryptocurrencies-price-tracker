package constant

import "time"

const (
	Port = ":8080"

	SQLiteDBName = "./cryptocurrencies.db"

	CoinCapHost        = "https://api.coincap.io"
	ExchangeRateHost   = "https://api.apilayer.com/exchangerates_data"
	ExchangeRateAPIKey = "CgWMWUg82yzKxdiuZ6UYFJWFrUKjzllZ"

	RedisHost       = "localhost:6379"
	RedisPass       = ""
	RedisExpiration = 15 * time.Minute

	PrefixKeyTokenRedis = "jwt-token-"
	JwtSecret           = "cryptocurrencies-price-tracker-secret"
)
