package myzo

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/url"
	"time"
)

type PotResponse struct {
	Pots []Pot `json:"pots"`
}

type Pot struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Style    string    `json:"style"`
	Balance  int64     `json:"balance"`
	Currency string    `json:"currency"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Deleted  bool      `json:"deleted"`
}

func basePotRequest(auth *Myzo, accountId string) *PotResponse {
	r, _ := auth.potResponseHandler(accountId)
	return r
}

func (auth *Myzo) GetAllPots(accountId string) []Pot {
	return basePotRequest(auth, accountId).Pots
}

// Amount is a Base64 integer, e.g. 100 = £/$/€1
func (auth *Myzo) DepositIntoPot(potId, accountId, amount string) (*PotResponse, error) {
	return auth.potMoveHandler(potId, accountId, amount, "deposit")
}

func (auth *Myzo) WithdrawFromPot(potId, accountId, amount string) (*PotResponse, error) {
	return auth.potMoveHandler(potId, accountId, amount, "withdraw")
}

func (auth *Myzo) potMoveHandler(potId, accountId, amount, event string) (*PotResponse, error) {
	data := url.Values{}
	if event == "deposit" {
		data.Set("source_account_id", accountId)
	} else if event == "withdraw" {
		data.Set("destination_account_id", accountId)
	} else {
		return nil, errors.New("invalid account")
	}
	data.Set("amount", amount)
	data.Set("dedupe_id", uuid.New().String())
	respBody, err := auth.postOrPutToMonzo("PUT", PotsEndpoint, "/"+potId+"/"+event, data)
	if err != nil {
		return nil, err
	}
	var potResponse PotResponse
	if err := json.Unmarshal(respBody, &potResponse); err != nil {
		return nil, err
	}
	return &potResponse, nil

}
