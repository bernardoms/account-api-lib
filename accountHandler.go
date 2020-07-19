package account_api_lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type AccountHandler struct {
	restClient HTTPClient
	Config     *AccountHandlerConfig
}

type Account struct {
	Data Data `json:"data"`
}

type Accounts struct {
	Data []Data `json:"data"`
}

type Data struct {
	Type           string     `json:"type"`
	Id             string     `json:"id"`
	OrganizationId string     `json:"organisation_id"`
	Version        int        `json:"version"`
	Attributes     Attributes `json:"attributes"`
}

type Attributes struct {
	Country               string `json:"country"`
	BaseCurrency          string `json:"base_currency"`
	AccountNumber         string `json:"account_number"`
	BankId                string `json:"bank_id"`
	BankIdCode            string `json:"bank_id_code"`
	CustomerId            string `json:"customer_id"`
	Bic                   string `json:"bic"`
	Iban                  string `json:"iban"`
	AccountClassification string `json:"Personal"`
	JointAccount          bool   `json:"joint_account"`
	Switched              bool   `json:"switched"`
	AccountMatchingOptOut bool   `json:"account_matching_opt_out"`
	Status                string `json:"status"`
}

func NewAccountHandler(client HTTPClient, config *AccountHandlerConfig) *AccountHandler {
	accountHandler := new(AccountHandler)
	accountHandler.restClient = client
	accountHandler.Config = config
	return accountHandler
}

func (ah *AccountHandler) CreateAccount(account *Account) (*Account, error) {
	j, err := json.Marshal(account)

	accountResponse := new(Account)

	req, err := http.NewRequest(http.MethodPost, ah.Config.URI, bytes.NewBuffer(j))

	var body []byte

	if err != nil {
		return nil, err
	}

	res, err := ah.restClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		body, err = ioutil.ReadAll(res.Body)

		err = json.Unmarshal(body, accountResponse)

		if err != nil {
			return nil, err
		}
	}

	if res.StatusCode >= http.StatusBadRequest && res.StatusCode < http.StatusInternalServerError {
		return nil, errors.New("client error " + string(body))
	}

	return accountResponse, err
}

func (ah *AccountHandler) FetchAccount(id string) (*Account, error) {
	req, err := http.NewRequest(http.MethodGet, ah.Config.URI+"/"+id, nil)

	resp, err := ah.restClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("id " + id + " not found")
	}

	account := new(Account)

	err = json.NewDecoder(resp.Body).Decode(account)

	return account, err
}

func (ah *AccountHandler) ListAccounts(pageNumber string, pageSize string) (*Accounts, error) {
	req, err := http.NewRequest(http.MethodGet, ah.Config.URI, nil)

	if err != nil {
		return nil, err
	}

	accounts := new(Accounts)

	q := req.URL.Query()
	q.Add("page[size]", pageSize)
	q.Add("page[number]", pageNumber)
	req.URL.RawQuery = q.Encode()

	resp, err := ah.restClient.Do(req)

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(accounts)

	return accounts, err
}

func (ah *AccountHandler) DeleteAccount(id string, version string) error {
	req, err := http.NewRequest(http.MethodDelete, ah.Config.URI+"/"+id, nil)

	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("version", version)

	req.URL.RawQuery = q.Encode()

	_, err = ah.restClient.Do(req)

	return err
}
