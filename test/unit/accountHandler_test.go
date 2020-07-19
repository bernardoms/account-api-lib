package unit

import (
	"bytes"
	"encoding/json"
	"errors"
	account_api_lib "github.com/bernardoms/account-api-lib"
	"github.com/bernardoms/account-api-lib/test/unit/mock"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_Create_account_success(t *testing.T) {
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("http://test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	account := new(account_api_lib.Account)

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

	response := new(http.Response)

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodPost })).Return(response, nil)

	_, err := accountHandler.CreateAccount(account)

	assert.Nil(t, err)
}

func Test_Create_account_4xx_error(t *testing.T) {
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("http://test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	account := new(account_api_lib.Account)

	response := new(http.Response)
	response.StatusCode = 400

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodPost })).Return(response, nil)

	_, err := accountHandler.CreateAccount(account)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "client error ")
}

func Test_Create_account_error_do_request(t *testing.T) {
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("http://test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	account := new(account_api_lib.Account)

	response := new(http.Response)

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodPost })).Return(response,
		errors.New("request error"))

	_, err := accountHandler.CreateAccount(account)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "request error")

}

func Test_Create_account_error_unmarshal_to_account(t *testing.T) {
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	account := new(account_api_lib.Account)

	response := new(http.Response)
	response.Body = ioutil.NopCloser(bytes.NewBufferString("wrong json format"))

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodPost })).Return(response,
		nil)

	_, err := accountHandler.CreateAccount(account)

	assert.NotNil(t, err)
}

func Test_Fetch_account_error_unmarshal_to_account(t *testing.T) {
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	response := new(http.Response)
	response.Body = ioutil.NopCloser(bytes.NewBufferString("wrong json format"))

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodGet })).Return(response,
		nil)

	_, err := accountHandler.FetchAccount("1")

	assert.NotNil(t, err)
}

func Test_Fetch_account_error_do_request(t *testing.T) {
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	response := new(http.Response)

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodGet })).Return(response,
		errors.New("request error"))

	_, err := accountHandler.FetchAccount("1")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "request error")
}

func Test_Fetch_account_error_not_found(t *testing.T) {
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	response := new(http.Response)
	response.StatusCode = 404

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodGet })).Return(response,
		nil)

	_, err := accountHandler.FetchAccount("1")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "id 1 not found")
}

func Test_Fetch_account_success(t *testing.T) {
	restClientMock := new(mock.RestClientMock)

	j, _ := json.Marshal(account_api_lib.Account{Data: account_api_lib.Data{Id: "1"}})

	config := account_api_lib.NewAccountHandleConfig("test.com")
	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	response := new(http.Response)
	response.Body = ioutil.NopCloser(bytes.NewBufferString(string(j)))

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodGet })).Return(response,
		nil)

	responseAccount, err := accountHandler.FetchAccount("1")

	assert.Nil(t, err)
	assert.NotNil(t, responseAccount.Data.Id)
}

func Test_List_account_error_unmarshal_to_account(t *testing.T) {
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	response := new(http.Response)
	response.Body = ioutil.NopCloser(bytes.NewBufferString("wrong json format"))

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodGet })).Return(response,
		nil)

	_, err := accountHandler.ListAccounts("0", "1")

	assert.NotNil(t, err)
}

func Test_List_account_error_do_request(t *testing.T) {
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	response := new(http.Response)

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodGet })).Return(response,
		errors.New("request error"))

	_, err := accountHandler.ListAccounts("0", "1")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "request error")
}

func Test_List_account_success(t *testing.T) {
	j, _ := json.Marshal(account_api_lib.Accounts{Data: nil})
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	response := new(http.Response)
	response.Body = ioutil.NopCloser(bytes.NewBufferString(string(j)))

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodGet })).Return(response, nil)

	responseAccounts, err := accountHandler.ListAccounts("0", "1")

	assert.Nil(t, err)
	assert.NotNil(t, responseAccounts)
}

func Test_Delete_account_error_do_request(t *testing.T) {
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	response := new(http.Response)

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodDelete })).Return(response,
		errors.New("request error"))

	err := accountHandler.DeleteAccount("0", "1")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "request error")
}

func Test_Delete_account_success(t *testing.T) {
	j, _ := json.Marshal(account_api_lib.Accounts{Data: nil})
	restClientMock := new(mock.RestClientMock)
	config := account_api_lib.NewAccountHandleConfig("test.com")

	accountHandler := account_api_lib.NewAccountHandler(restClientMock, config)

	response := new(http.Response)
	response.Body = ioutil.NopCloser(bytes.NewBufferString(string(j)))

	restClientMock.On("Do",
		mock2.MatchedBy(func(req *http.Request) bool { return req != nil && req.Method == http.MethodDelete })).Return(response, nil)

	err := accountHandler.DeleteAccount("1", "0")

	assert.Nil(t, err)
}
