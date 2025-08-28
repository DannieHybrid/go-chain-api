Go Chain API

A simple Go + Gin API to fetch Ethereum account balances and ERC20 token balances.

🚀 Features

Get native ETH balance of any Ethereum address

Get ERC20 token balance for any Ethereum address

Built with Go
, Gin
, and go-ethereum

📦 Installation

Clone the repo:

git clone https://github.com/YOUR_USERNAME/go-chain-api.git
cd go-chain-api


Install dependencies:

go mod tidy

⚡ Usage

Run the API:

go run main.go


Server will start on http://localhost:8080

Endpoints
1. Get ETH Balance
GET /balance/:address


Example:

curl http://localhost:8080/balance/0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045


Response:

{
  "address": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
  "balance": "4787485422738848697"
}

2. Get ERC20 Token Balance
GET /token/:contract/:address


Example (USDT balance):

curl http://localhost:8080/token/0xdAC17F958D2ee523a2206206994597C13D831ec7/0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045


Response:

{
  "contract": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
  "address": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
  "balance": "123456789"
}

🛠️ Tech Stack

Go

Gin Framework

go-ethereum

📄 License

MIT
