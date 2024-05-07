# CRYPTOCURRENCIES PRICE TRACKER

This is a Cryptocurrencies price tracker application providing a REST API to show the price of user tracked coins. 

This application uses SQLite for its database. It utilizes the client `https://docs.coincap.io` as a source of price information and `https://exchangeratesapi.io/documentation/` as a source of exchange rate data, which is stored in cache.

`config.go` is a configuration file that contains values such as host, port, key, secret, and others used by the application.


# REST API

## Sign Up/Register User

This is a POST request, for create/register user by submitting data to an API via the request body. A successful POST request typically returns a 200 OK

### Request

`POST /user`

    curl --location 'http://localhost:8080/user' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "Mickey Mouse",
        "email": "mickey.mouse@disney.com",
        "username": "mickeymouse",
        "password": "Mickey123!",
        "confirmPassword": "Mickey123!"
    }'

### Response

    HTTP/1.1 200 OK
    {
        "responseMessage": "Success"
    }

## Login User

The Login API is used to authenticate a user in this application. Upon successful login, the user will receive a token to access all endpoints that available in the system.

### Request

`POST /user/login`

    curl --location 'http://localhost:8080/user/login' \
    --header 'Content-Type: application/json' \
    --data '{
        "username": "mickeymouse",
        "password": "Mickey123!"
    }'

    
### Response

    HTTP/1.1 200 OK
    {
        "responseMessage": "Success",
        "data": {
            "name": "Mickey Mouse",
            "username": "mickeymouse",
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjMsInVzZXJuYW1lIjoibWlja2V5bW91c2UiLCJleHAiOjE3MTUwNDkxMDd9.jOxue-Yy6rNKLe7Ab1f1XqRReZcMNWDWOdEaV4vGeqg"
        }
    }

## Logout User

The Logout API is used to invalidate the token of a logged-in user, thereby effectively logging them out of the system and revoking their access to protected endpoints.

### Request

`POST /user/logout`

    curl --location --request POST 'http://localhost:8080/user/logout' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjMsInVzZXJuYW1lIjoibWlja2V5bW91c2UiLCJleHAiOjE3MTUwNDkxMDd9.jOxue-Yy6rNKLe7Ab1f1XqRReZcMNWDWOdEaV4vGeqg' \
    --data ''

### Response

    HTTP/1.1 200 OK
    {
        "responseMessage": "Success"
    }

## Asset List

This is a GET request for get all cryptocurrency assets and their price in USD.

### Request

`GET /assets`

    curl --location --request GET 'http://localhost:8080/assets' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjMsInVzZXJuYW1lIjoibWlja2V5bW91c2UiLCJleHAiOjE3MTUwNDk3Nzd9.b23t4HGF1481vQDyLwMUPUBgm0XZMZxUttnSWE-TXTg' \
    --data '{}'

### Response

    HTTP/1.1 200 OK
    {
        "responseMessage": "Success",
        "data": {
            "data": [
                {
                    "id": "bitcoin",
                    "rank": "1",
                    "symbol": "BTC",
                    "name": "Bitcoin",
                    "supply": "19694909.0000000000000000",
                    "maxSupply": "21000000.0000000000000000",
                    "marketCapUsd": "1256694752864.1173169154847194",
                    "volumeUsd24Hr": "8260313600.5021313343131909",
                    "priceUsd": "63808.1015182206384866",
                    "changePercent24Hr": "-0.5980805214553729",
                    "vwap24Hr": "63935.8699751004696171"
                },
                {
                    "id": "ethereum",
                    "rank": "2",
                    "symbol": "ETH",
                    "name": "Ethereum",
                    "supply": "120099596.6021065300000000",
                    "maxSupply": "",
                    "marketCapUsd": "371506963742.4482159951479142",
                    "volumeUsd24Hr": "5550054570.5986830247567374",
                    "priceUsd": "3093.3239931959274891",
                    "changePercent24Hr": "-1.7995774800281209",
                    "vwap24Hr": "3124.4907090270044477"
                },
            ...
            ]
        }
    }

## Get User Tracked Coins

This is a GET endpoint to show user list of tracked coins and their price in IDR.

### Request

`GET /user/tracker`

    curl --location --request GET 'http://localhost:8080/user/tracker' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjMsInVzZXJuYW1lIjoibWlja2V5bW91c2UiLCJleHAiOjE3MTUwNDk3Nzd9.b23t4HGF1481vQDyLwMUPUBgm0XZMZxUttnSWE-TXTg' \
    --data '{}'

### Response

    HTTP/1.1 200 OK
    {
        "responseMessage": "Success",
        "data": {
            "username": "mickeymouse",
            "TrackedCoins": [
                {
                    "coinId": "bitcoin-cash",
                    "coinSymbol": "BCH",
                    "coinName": "Bitcoin Cash",
                    "priceIDR": 7728091.827896823
                },
                {
                    "coinId": "bitcoin",
                    "coinSymbol": "BTC",
                    "coinName": "Bitcoin",
                    "priceIDR": 1023592465.6179231
                },
                {
                    "coinId": "ethereum",
                    "coinSymbol": "ETH",
                    "coinName": "Ethereum",
                    "priceIDR": 49622957.209015705
                }
            ]
        }
    }

## Add User Tracked Coin

This is a POST endpoint for user can add coin to tracker.

### Request

`POST /user/tracker`

    curl --location 'http://localhost:8080/user/tracker' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjMsInVzZXJuYW1lIjoibWlja2V5bW91c2UiLCJleHAiOjE3MTUwNDk3Nzd9.b23t4HGF1481vQDyLwMUPUBgm0XZMZxUttnSWE-TXTg' \
    --data '{
        "coinId": "ethereum"
    }'

### Response

    HTTP/1.1 200 OK
    {
        "responseMessage": "Success"
    }

## Delete User Tracked Coin

This is a DELETE endpoint for user can remove coin from tracker.

### Request

`DELETE /user/tracker/{coinId}`

    curl --location --request DELETE 'http://localhost:8080/user/tracker/bitcoin' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjMsInVzZXJuYW1lIjoibWlja2V5bW91c2UiLCJleHAiOjE3MTUwNTA3MjR9.2DUWIJTQ8ALZCuGUM_FRqOGTtlYaBP0QgGDzI0ig7ZM' \
    --data ''

### Response

    HTTP/1.1 200 OK
    {
        "responseMessage": "Success"
    }
