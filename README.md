# go-chain-api

A lightweight Ethereum API service written in Go, inspired by Zerion’s mission of powering on-chain finance.  
It provides clean REST endpoints for fetching balances and token holdings directly from the Ethereum blockchain.

---

## 🚀 Features
- **ETH Balance Endpoint** – Get ETH balance of any valid Ethereum address.  
- **ERC20 Token Balance Endpoint** – Query balances of ERC20 tokens by contract + wallet.  
- **Input Validation & Error Handling** – Clear errors for invalid addresses and failed contract calls.  
- **Go-Powered Backend** – Built with Go’s standard library and go-ethereum bindings for speed and reliability.  
- **Production-Ready Structure** – Modular codebase with simple startup flow and API router.

---

## 📦 Installation

Clone the repo and install dependencies:

```bash
git clone https://github.com/DannieHybrid/go-chain-api.git
cd go-chain-api
go mod tidy
⚡ Running the Service
Start the API server:

bash
Copy
Edit
go run main.go
Server will start on http://localhost:8080

🔗 API Endpoints
1. Get ETH Balance
bash
Copy
Edit
GET /balance/{wallet_address}
Example:

bash
Copy
Edit
curl http://localhost:8080/balance/0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045
Response:

json
Copy
Edit
{
  "address": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
  "balance": "4787485422738848697"
}
2. Get ERC20 Token Balance
bash
Copy
Edit
GET /token/{token_contract}/{wallet_address}
Example:

bash
Copy
Edit
curl http://localhost:8080/token/0xdAC17F958D2ee523a2206206994597C13D831ec7/0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045
Response:

json
Copy
Edit
{
  "token": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
  "address": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
  "balance": "123456789"
}
🛠 Tech Stack
Go (API backend, routing, and execution)

go-ethereum (github.com/ethereum/go-ethereum)

JSON REST API for easy integration


