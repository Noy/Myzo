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
	ID               string `json:"id"`
	Name             string `json:"name"`
	Style            string `json:"style"`
	Balance          int64  `json:"balance"`
	Currency         string `json:"currency"`
	GoalAmount       int64  `json:"goal_amount"`
	Type             string `json:"type"`
	ProductId        string `json:"product_id"`
	CurrentAccountId string `json:"current_account_id"`
	CoverImageUrl    string `json:"cover_image_url"`
	IsaWrapper       string `json:"isa_wrapper"`
	RoundUp          bool   `json:"round_up"`
	//RoundUpMultiplier int64?
	IsTaxPot          bool      `json:"is_tax_pot"` // FOR BUSINESSES
	Created           time.Time `json:"created"`
	Updated           time.Time `json:"updated"`
	Deleted           bool      `json:"deleted"`
	Locked            bool      `json:"locked"`
	CharityId         string    `json:"charity_id"`
	AvailableForBills bool      `json:"available_for_bills"`

	//Friendly
	BalanceFriendly    float64
	GoalAmountFriendly float64
	CreatedFriendly    string
	UpdatedFriendly    string
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
