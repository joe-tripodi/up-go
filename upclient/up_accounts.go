package upclient

import (
	"encoding/json"
	"log"
)

type AccountResponse struct {
	Data  []Account `json:"data"`
	Links Links     `json:"links"`
}

type Links struct {
	Prev string `json:"prev"`
	Next string `json:"next"`
}

type Attributes struct {
	DisplayName   string  `json:"displayName"`
	AccountType   string  `json:"accountType"`
	OwnershipType string  `json:"ownershipType"`
	Balance       Balance `json:"balance"`
	CreatedAt     string  `json:"createdAt"`
}

type Balance struct {
	CurrencyCode     string `json:"currencyCode"`
	Value            string `json:"value"`
	ValueInBaseUnits int    `json:"valueInBaseUnits"`
}

type Account struct {
	Type       string     `json:"type"`
	Id         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

func (up *UpClient) GetAccounts() ([]Account, error) {
	data, err := up.doRequest("GET", "/accounts", nil)
	if err != nil {
		log.Printf("ERROR: Failed to GET /accounts: %v", err)
		return nil, err
	}

	var accountResponse AccountResponse
	if err = json.Unmarshal(data, &accountResponse); err != nil {
		log.Printf("ERROR: Failed to Unmarshal account data: %v", err)
		return nil, err
	}

	return accountResponse.Data, nil
}

func (accounts Account) Print() error {
	return nil
}
