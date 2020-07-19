package integrated

import (
	"bytes"
	"encoding/json"
	accountapilib "github.com/bernardoms/account-api-lib"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	account1 := new(accountapilib.Account)

	account1.Data.Id = "123e4567-e89b-12d3-a456-426655440001"
	account1.Data.OrganizationId = "123e4567-e89b-12d3-a456-426655440001"
	account1.Data.Type = "accounts"
	account1.Data.Attributes.Country = "GB"
	account1.Data.Attributes.BaseCurrency = "GBP"
	account1.Data.Attributes.BankId = "123"
	account1.Data.Attributes.BankIdCode = "GBDSC"
	account1.Data.Attributes.AccountNumber = "10000004"
	account1.Data.Attributes.CustomerId = "24123"
	account1.Data.Attributes.Bic = "NWBKGB42"
	account1.Data.Attributes.Iban = "GB28NWBK40030212764204"
	account1.Data.Attributes.AccountClassification = "346345"

	j, _ := json.Marshal(account1)

	r, _ := http.NewRequest(http.MethodPost, GetUrl(), bytes.NewBuffer(j))
	_, _ = http.DefaultClient.Do(r)

	account2 := new(accountapilib.Account)

	account2.Data.Id = "123e4567-e89b-12d3-a456-426655440002"
	account2.Data.OrganizationId = "123e4567-e89b-12d3-a456-426655440001"
	account2.Data.Type = "accounts"
	account2.Data.Attributes.Country = "GB"
	account2.Data.Attributes.BaseCurrency = "GBP"
	account2.Data.Attributes.BankId = "123"
	account2.Data.Attributes.BankIdCode = "GBDSC"
	account2.Data.Attributes.AccountNumber = "10000004"
	account2.Data.Attributes.CustomerId = "24123"
	account2.Data.Attributes.Bic = "NWBKGB42"
	account2.Data.Attributes.Iban = "GB28NWBK40030212764204"
	account2.Data.Attributes.AccountClassification = "346345"

	j, _ = json.Marshal(account2)

	r, _ = http.NewRequest(http.MethodPost, GetUrl(), bytes.NewBuffer(j))
	_, _ = http.DefaultClient.Do(r)

	os.Exit(m.Run())
}

func GetUrl() string {
	if os.Getenv("TEST_URL") != "" {
		return os.Getenv("TEST_URL")
	}
	return "http://localhost:8080/v1/organisation/accounts"
}
