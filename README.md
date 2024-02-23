# Cashfree Payout Go SDK
![GitHub](https://img.shields.io/github/license/cashfree/cashfree-payout) ![Discord](https://img.shields.io/discord/931125665669972018?label=discord) ![GitHub last commit (branch)](https://img.shields.io/github/last-commit/cashfree/cashfree-pg/main) ![GitHub release (with filter)](https://img.shields.io/github/v/release/cashfree/cashfree-payout?label=latest) ![GitHub forks](https://img.shields.io/github/forks/cashfree/cashfree-payout) [![Coverage Status](https://coveralls.io/repos/github/cashfree/cashfree-pg/badge.svg?branch=main)](https://coveralls.io/github/cashfree/cashfree-pg?branch=main) [![GoDoc](https://godoc.org/github.com/cashfree/cashfree-pg/v4?status.svg)](https://godoc.org/github.com/cashfree/cashfree-pg/v4)

The Cashfree Payout Go SDK offers a convenient solution to access [Cashfree Payout APIs](https://docs.cashfree.com/reference/pg-new-apis-endpoint) from a server-side Go  applications.



## Documentation

Cashfree's Payout API Documentation - //todo add link of documentation

Learn and understand payment gateway workflows at Cashfree Payments [here](https://docs.cashfree.com/docs/payouts)

Try out our interactive guides at [Cashfree Dev Studio](https://www.cashfree.com/devstudio) !

## Getting Started

### Installation
```bash
go get github.com/cashfree/cashfree-payout
```
### Configuration

```go 
import (
    cashfree "github.com/cashfree/cashfree-payout"
)

clientId := "<x-client-id>"
clientSecret := "<x-client-secret>"
cashfree.XClientId = &clientId
cashfree.XClientSecret = &clientSecret
cashfree.XEnvironment = cashfree.SANDBOX
```

Generate your API keys (x-client-id , x-client-secret) from [Cashfree Merchant Dashboard](https://merchant.cashfree.com/merchants/login)


## Supported Resources

- [Beneficiary](docs/BeneficiaryV2API.md)

- [Transfer](docs/TransfersV2API.md)


## Licence

Apache Licensed. See [LICENSE.md](LICENSE.md) for more details
