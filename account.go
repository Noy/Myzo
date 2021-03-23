package myzo

import "time"

type AccountResponse struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Name          string    `json:"name"`
	ID            string    `json:"id"`
	Description   string    `json:"description"`
	Created       time.Time `json:"created"`
	Closed        bool      `json:"closed"`
	Type          string    `json:"type"`
	Currency      string    `json:"currency"`
	CountryCode   string    `json:"country_code"`
	Owners        []Owner   `json:"owners"`
	AccountNumber string    `json:"account_number"`
	SortCode      string    `json:"sort_code"`
}

type Owner struct {
	UserID             string `json:"user_id"`
	PreferredName      string `json:"preferred_name"`
	PreferredFirstName string `json:"preferred_first_name"`
}

func baseAccountRequestRequest(auth *Myzo) *AccountResponse {
	r, _ := auth.accountResponseHandler()
	return r
}

func (auth *Myzo) GetAccounts() []Account {
	return baseAccountRequestRequest(auth).Accounts
}

func (auth *Myzo) WhoAmI() (string, error) {
	body, err := auth.getFromMonzo("/ping/whoami", "", "")
	if err != nil {
		return "", err
	} else {
		return string(body), nil
	}
}

func (auth *Myzo) VerifiedThroughApp() (string, error) {
	body, err := auth.getFromMonzo("/accounts", "", "")
	if err != nil {
		return "", err
	} else {
		return string(body), nil
	}
}
