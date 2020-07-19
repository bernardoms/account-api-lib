# Account-api-lib - Bernardo Monteiro da Silva

## Go Experience

* I have 6 months working with go lang, I use go lang to write some lambdas on AWS.

## Installation

```
go get github.com/bernardoms/account-api-lib
```
## How to use
* For creating a new account ex:
```go
package account_api_lib

import (
accountapilib "github.com/bernardoms/account-api-lib"
"net/http"
)
func createAccount() {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig("http://localhost:8080")

	accountHandler := accountapilib.NewAccountHandler(client, config)

	account := new(accountapilib.Account)

	account.Data.Id = "123e4567-e89b-12d3-a456-426655440000"
	account.Data.OrganizationId = "123e4567-e89b-12d3-a456-426655440001"
	account.Data.Type = "accounts"
	account.Data.Attributes.Country = "GB"
	account.Data.Attributes.BaseCurrency = "GBP"
	account.Data.Attributes.BankId = "123"
	account.Data.Attributes.BankIdCode = "GBDSC"
	account.Data.Attributes.AccountNumber = "10000004"
	account.Data.Attributes.CustomerId = "24123"
	account.Data.Attributes.Bic = "NWBKGB42"
	account.Data.Attributes.Iban = "GB28NWBK40030212764204"
	account.Data.Attributes.AccountClassification = "346345"

	accountHandler.CreateAccount(account)
}
```

* For deleting an account ex:

```go
package account_api_lib

import (
accountapilib "github.com/bernardoms/account-api-lib"
"net/http"
)
func deleteAccount() {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig("http://localhost:8080")

	accountHandler := accountapilib.NewAccountHandler(client, config)
    
    accountHandler.DeleteAccount("123e4567-e89b-12d3-a456-426655440000", "0")
}
```

* For Fetch an account ex: 

```go
package account_api_lib

import (
accountapilib "github.com/bernardoms/account-api-lib"
"net/http"
)

func fetchAccount() {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig("http://localhost:8080")

	accountHandler := accountapilib.NewAccountHandler(client, config)
    
    accountHandler.FetchAccount("123e4567-e89b-12d3-a456-426655440000")
}
```
 
* For List accounts ex:

```go
package account_api_lib

import (
accountapilib "github.com/bernardoms/account-api-lib"
"net/http"
)

func fetchAccount() {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig("http://localhost:8080")

	accountHandler := accountapilib.NewAccountHandler(client, config)
    
    accountHandler.ListAccounts("0", "100")
}
```

## Testing
It's possible to test using the makefile on the project

* unit-test:
`make unit-test`

* integration-test:
`make integration-test`
 
It's possible to test using the docker-compose command
`docker-compose up`

## Decisions
* Decided to use go module, because a lot easier than using go path with go get.

* Decided to create a make file for be easier to run integration test without a lot of effort or using a lib 
for bring up a container on test.

* Using interface for http client (easier to test).

* Used dependency injection because It's made life easier for inject mock for test or inject 
your own client implementation.
