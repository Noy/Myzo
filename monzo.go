package myzo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	BaseURL = "https://api.monzo.com"
	BalanceEndpoint = "/balance"
	TransactionsEndpoint = "/transactions"
	AccountEndpoint = "/accounts"
	PotsEndpoint = "/pots"
)

type Myzo struct {
	ClientID string
	UserID string
	AccessToken string
	AccountID string

	Debug bool
	ResponseBody []byte
}

/**
	Authentication with Monzo.
*/
func (auth *Myzo) authenticate(method, endpoint, params string) ([]byte, error) {
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

/**
	Base request for handling balance responses.
*/
func (auth *Myzo) balanceResponseHandler() (*BalanceResponse, error) {
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

/**
	Base request for handling Pot responses.
*/
func (auth *Myzo) potResponseHandler() (*PotResponse, error) {
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

/**
	Base request for handling transaction responses.
*/
func (auth *Myzo) transactionResponseHandler(bulkRequest bool, daysAgo int, expandBy, optionalId string) (*TransactionsResponse, error) {
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

/**
	Base request for handling account responses.
 */
func (auth *Myzo) accountResponseHandler() (*AccountResponse, error) {
	resp, err := auth.authenticate("GET", AccountEndpoint,"")
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
func (auth *Myzo) FeedHandler(URL string, params []string) ([]byte, error) {
	client := &http.Client{}
	data := url.Values{}
	if len(params) != 6 {
		log.Fatal("This needs exactly 6 parameters, no more, no less.")
		return nil, nil
	}
	data.Set("params[title]", params[0])
	data.Set("params[image_url]", params[1])
	data.Set("params[background_color]", params[2])
	data.Set("params[body_color]", params[3])
	data.Set("params[title_color]", params[4])
	data.Set("params[body]", params[5])
	req, err := http.NewRequest("POST",
		BaseURL+"/feed?account_id="+auth.AccountID+"&type=basic"+"&url="+URL,
		bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+auth.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if err != nil {
		log.Println(err.Error())
	}
	defer req.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}