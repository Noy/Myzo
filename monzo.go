package monzo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	BaseURL = "https://api.monzo.com"
	BalanceEndpoint = "/balance"
	TransactionsEndpoint = "/transactions"
	PotsEndpoint = "/pots"
)

type AuthMonzo struct {
	ClientID string
	UserID string
	AccessToken string
	AccountID string

	Debug bool
	ResponseBody []byte
}

func (auth *AuthMonzo) authenticate(method, endpoint, params string) ([]byte, error) {
	client := &http.Client{Timeout:time.Second * 2}
	req, err := http.NewRequest(method, BaseURL+endpoint+"?account_id="+auth.AccountID+params, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+auth.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResponse, err := ioutil.ReadAll(resp.Body)
	if auth.Debug {
		auth.ResponseBody = jsonResponse
		log.Println(string(jsonResponse))
	}
	if err != nil {
		return nil, err
	}
	return jsonResponse, nil
}

func (auth *AuthMonzo) balanceResponseHandler() (*BalanceResponse, error) {
	resp, err := auth.authenticate("GET", BalanceEndpoint,  "")
	if err != nil {
		return nil, err
	}
	var balanceResponse BalanceResponse
	if err := json.Unmarshal(resp, &balanceResponse); err != nil {
		return nil, err
	}
	return &balanceResponse, nil
}

func (auth *AuthMonzo) potResponseHandler() (*PotResponse, error) {
	resp, err := auth.authenticate("GET", PotsEndpoint, "")
	if err != nil {
		return nil, err
	}
	var potResponse PotResponse
	if err := json.Unmarshal(resp, &potResponse); err != nil {
		return nil, err
	}
	return &potResponse, nil
}

func (auth *AuthMonzo) transactionResponseHandler(bulkRequest bool, daysAgo int, expandBy, optionalId string) (*TransactionsResponse, error) {
	split := strings.Split(time.Now().AddDate(0,0, -daysAgo).Format(time.RFC3339), "+")
	var resp []byte
	var err error
	if bulkRequest {
		resp, err = auth.authenticate("GET", TransactionsEndpoint+optionalId, "&since="+split[0]+"Z&expand[]="+expandBy)
	} else {
		resp, err = auth.authenticate("GET", TransactionsEndpoint+optionalId, "&expand[]="+expandBy)
	}
	if err != nil {
		return nil, err
	}
	var transactionResponse TransactionsResponse
	if err := json.Unmarshal(resp, &transactionResponse); err != nil {
		return nil, err
	}
	return &transactionResponse, nil
}
