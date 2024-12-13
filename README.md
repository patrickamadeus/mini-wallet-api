# Mini Wallet API

This project is a simple wallet API built using Go (Golang) and the Gin web framework. The API allows users to initialize accounts, enable and disable wallets, view wallet information, and perform transactions such as deposits and withdrawals.

## **Table of Contents**
- [Mini Wallet API](#mini-wallet-api)
  - [**Table of Contents**](#table-of-contents)
  - [**Features**](#features)
  - [**Project Structure**](#project-structure)
  - [**Prerequisites**](#prerequisites)
  - [**Setup Instructions**](#setup-instructions)
  - [**API Documentation**](#api-documentation)

---

## **Features**
- Account initialization
- Wallet activation and deactivation
- Wallet balance view
- Deposit and withdrawal of funds
- Transaction history view

---

## **Project Structure**
```
mini-wallet-api/
├── README.md
├── go.mod
├── go.sum
├── main.go                     # Entrypoint of app
├── handlers                    # API route handlers
│   ├── deposit.go
│   ├── disable.go              
│   ├── enable.go
│   ├── init.go
│   ├── view_transaction.go
│   ├── view_wallet.go
│   └── withdraw.go
├── middleware                  # Auth Middleware
│   └── auth.go
├── models                      # Data models
│   ├── transaction.go
│   └── wallet.go
└── utils                       # Utility function
    ├── jwt.go
    └── utils.go
```

---

## **Prerequisites**
- **Go** (version 1.23 or higher) installed on your system
- **Git** to clone the repository

---

## **Setup Instructions**
1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd mini-wallet-api
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the server**:
   ```bash
   go run main.go
   ```

4. **Access the API** at:
   ```
   http://localhost:8080
   ```

---

## **API Documentation**

For more detail please refer to this [Postman Documentary](https://documenter.getpostman.com/view/8411283/SVfMSqA3?version=latest).

---

With this documentation, you should be able to set up, run, and use the Mini Wallet API with ease. If you have any questions or issues, please contact patrickai0309@gmail.com / open an issue.

