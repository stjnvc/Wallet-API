# Wallet API

The Wallet API is a JSON API written in Golang that manages the wallets of players in an online casino. It provides endpoints for retrieving wallet balances and performing credit and debit operations on wallets. The API is designed to be flexible and easily testable, with support for multiple storage mechanisms.

## Endpoints

- **Balance**: Retrieves the balance of a given wallet ID.
  - `GET /api/v1/wallets/{wallet_id}/balance`
- **Credit**: Credits money on a given wallet ID.
  - `POST /api/v1/wallets/{wallet_id}/credit`
- **Debit**: Debits money from a given wallet ID.
  - `POST /api/v1/wallets/{wallet_id}/debit`
- **Login**: Login route to obtain jwt.
  - `POST /api/v1/auth/login`
- **Register**: Create dummy user (username, email, password) for authorization.
  - `POST /api/v1/auth/register`

## Business Rules

- A wallet balance cannot go below 0.
- Amounts sent in the credit and debit operations cannot be negative.

## Bonus Tasks

- Cache the wallet balances in Redis for faster retrieval.
- Implement an auth endpoint and authentication verification.
- Log incoming requests for monitoring and debugging.

## Libraries Used

- HTTP: [Gin](https://github.com/gin-gonic/gin)
- MySQL: [GORM](https://github.com/go-gorm/gorm)
- Redis: [go-redis](https://github.com/go-redis/redis)
- Numbers: [Decimal](https://github.com/shopspring/decimal)
- Logger: [Logrus](https://github.com/sirupsen/logrus)

## Setup

1. Clone the repository.
2. Install dependencies: `go mod tidy`.
3. Inside project root directory un docker compose up --build 
4. The API should now be accessible at `http://localhost:8080`.


## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
