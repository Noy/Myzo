package myzo

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	BaseURL              = "https://api.monzo.com"
	BalanceEndpoint      = "/balance"
	TransactionsEndpoint = "/transactions"
	AccountEndpoint      = "/accounts"
	PotsEndpoint         = "/pots"
)

type Myzo struct {
	ClientID    string
	UserID      string
	AccessToken string
	AccountIDs  map[string]string

	Debug        bool
	ResponseBody []byte
}

/**
Authentication with Monzo.
*/
func (auth *Myzo) authenticate(method, url string, data io.Reader) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + auth.AccessToken)
	if method == "PUT" || method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonResponse, err := ioutil.ReadAll(resp.Body)
	if auth.Debug {
		log.Println(url)
		auth.ResponseBody = jsonResponse
		log.Println(string(jsonResponse))
	}
	if err != nil {
		return nil, err
	}
	return jsonResponse, nil
}

/**
For all GET requests
*/
func (auth *Myzo) getFromMonzo(endpoint, params, accountId string) ([]byte, error) {
	var accountIdParam = "?account_id="
	if endpoint == PotsEndpoint {
		accountIdParam = "?current_account_id="
	}
	jsonResponse, err := auth.authenticate("GET", BaseURL+endpoint+accountIdParam+accountId+params, nil)
	if err != nil {
		return nil, err
	}
	return jsonResponse, nil
}

/**
For all POST/PUT requests, please note the authentication adds the 'application/x-www-form-urlencoded' for all of these requests.
*/
func (auth *Myzo) postOrPutToMonzo(method, endPoint, extraData string, data url.Values) ([]byte, error) {
	resp, err := auth.authenticate(method, BaseURL+endPoint+extraData, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/**
Base request for handling balance responses.
*/
func (auth *Myzo) balanceResponseHandler(accountId string) (*BalanceResponse, error) {
	resp, err := auth.getFromMonzo(BalanceEndpoint, "", accountId)
	if err != nil {
		return nil, err
	}
	var balanceResponse BalanceResponse
	if err := json.Unmarshal(resp, &balanceResponse); err != nil {
		return nil, err
	}
	return &balanceResponse, nil
}

/**
Base request for handling Pot responses.
*/
func (auth *Myzo) potResponseHandler(accountId string) (*PotResponse, error) {
	resp, err := auth.getFromMonzo(PotsEndpoint, "", accountId)
	if err != nil {
		return nil, err
	}
	var potResponse PotResponse
	if err := json.Unmarshal(resp, &potResponse); err != nil {
		return nil, err
	}
	return &potResponse, nil
}

/**
Base request for handling transaction responses.
*/
func (auth *Myzo) transactionResponseHandler(bulkRequest bool, daysAgo, before int, expandBy, optionalId, accountId string) (*TransactionsResponse, error) {
	split := strings.Split(time.Now().AddDate(0, 0, -daysAgo).Format(time.RFC3339), "+")
	splitBefore := strings.Split(time.Now().AddDate(0, 0, -before).Format(time.RFC3339), "+")
	var resp []byte
	var err error
	if bulkRequest {
		resp, err = auth.getFromMonzo(TransactionsEndpoint+optionalId, "&since="+split[0]+"&expand[]="+expandBy+"&before="+splitBefore[0], accountId)
	} else {
		resp, err = auth.getFromMonzo(TransactionsEndpoint+optionalId, "&expand[]="+expandBy, accountId)
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

/**
Base request for handling account responses.
*/
func (auth *Myzo) accountResponseHandler() (*AccountResponse, error) {
	resp, err := auth.getFromMonzo(AccountEndpoint, "", "") // no need for account ID
	if err != nil {
		return nil, err
	}
	var accountResponse AccountResponse
	if err := json.Unmarshal(resp, &accountResponse); err != nil {
		return nil, err
	}
	return &accountResponse, nil
}

/**
Send custom feed items to your timeline.
*/
func (auth *Myzo) FeedHandler(URL string, params [6]string, accountId string) ([]byte, error) {
	data := url.Values{}
	data.Set("params[title]", params[0])
	data.Set("params[image_url]", params[1])
	data.Set("params[background_color]", params[2])
	data.Set("params[body_color]", params[3])
	data.Set("params[title_color]", params[4])
	data.Set("params[body]", params[5])
	respBody, err := auth.postOrPutToMonzo("POST", "/feed", "?account_id="+accountId+"&type=basic"+"&url="+URL, data)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
