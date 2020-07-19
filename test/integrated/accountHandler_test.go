package integrated

import (
	accountapilib "github.com/bernardoms/account-api-lib"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_create_account_success(t *testing.T) {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig(GetUrl())

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

	_, err := accountHandler.CreateAccount(account)

	//Always delete the account after testing

	req, _ := http.NewRequest(http.MethodDelete, config.URI+"/"+account.Data.Id+"?version=0", nil)

	defer client.Do(req)

	assert.Nil(t, err)
}

func Test_create_account_4xx_error(t *testing.T) {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig(GetUrl())

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
	account.Data.Attributes.Bic = "123"
	account.Data.Attributes.Iban = "234"
	account.Data.Attributes.AccountClassification = "346345"

	_, err := accountHandler.CreateAccount(account)

	assert.NotNil(t, err)
}

func Test_Fetch_account_success(t *testing.T) {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig(GetUrl())

	accountHandler := accountapilib.NewAccountHandler(client, config)

	response, err := accountHandler.FetchAccount("123e4567-e89b-12d3-a456-426655440001")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426655440001", response.Data.Id)
}

func Test_Fetch_account_not_found(t *testing.T) {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig(GetUrl())

	accountHandler := accountapilib.NewAccountHandler(client, config)

	response, err := accountHandler.FetchAccount("123e4567-e89b-12d3-a456-426655440006")

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func Test_List_accounts_paging_one_account(t *testing.T) {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig(GetUrl())

	accountHandler := accountapilib.NewAccountHandler(client, config)

	response, err := accountHandler.ListAccounts("0", "1")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.Data))
}

func Test_List_accounts_paging_all_elements(t *testing.T) {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig(GetUrl())

	accountHandler := accountapilib.NewAccountHandler(client, config)

	response, err := accountHandler.ListAccounts("0", "100")

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 2, len(response.Data))
}

func Test_Delete_account_success(t *testing.T) {
	client := http.DefaultClient

	config := accountapilib.NewAccountHandleConfig(GetUrl())

	accountHandler := accountapilib.NewAccountHandler(client, config)

	err := accountHandler.DeleteAccount("0", "0")

	assert.Nil(t, err)
}
