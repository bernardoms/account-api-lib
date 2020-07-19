package account_api_lib

import "net/http"

func test() {
	client := http.DefaultClient

	config := NewAccountHandleConfig("http://localhost:8080")

	accountHandler := NewAccountHandler(client, config)

	account := new(Account)

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
