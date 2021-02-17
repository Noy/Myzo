package myzo

import "time"

type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
	Transaction  Transaction   `json:"transaction"`
}

type Transaction struct {
	AccountBalance             int64     `json:"account_balance"`
	Amount                     int64     `json:"amount"`
	Created                    time.Time `json:"created"`
	Currency                   string    `json:"currency"`
	Description                string    `json:"description"`
	ID                         string    `json:"id"`
	MetaData                   struct{}  `json:"metadata"`
	Notes                      string    `json:"notes"`
	IsLoad                     bool      `json:"is_load"`
	Settled                    string    `json:"settled"`
	Category                   string    `json:"category"`
	LocalAmount                int64     `json:"local_amount"`
	LocalCurrency              string    `json:"local_currency"`
	CanBeExcludedFromBreakdown bool      `json:"can_be_excluded_from_breakdown"`
	CanBeMadeSubscription      bool      `json:"can_be_made_subscription"`
	CanSplitBill               bool      `json:"can_split_the_bill"`
	CanAddToTab                bool      `json:"can_add_to_tab"`
	AmountIsPending            bool      `json:"amount_is_pending"`
	Originator                 bool      `json:"originator"`
	DedupeID                   string    `json:"dedupe_id"`
	Scheme                     string    `json:"scheme"`
	Merchant                   Merchant  `json:"merchant"`
	//Misc, if you want to parse the dates as a string instead of time.Time
	CreatedString string
	MerchantCreated string
}

type Merchant struct {
	ID       string    `json:"id"`
	GroupID  string    `json:"group_id"`
	Created  time.Time `json:"created"`
	Name     string    `json:"name"`
	Logo     string    `json:"logo"`
	Emoji    string    `json:"emoji"`
	Category string    `json:"category"`
	Online   bool      `json:"online"`
	ATM      bool      `json:"atm"`
	Address  struct {
		ShortFormatted string  `json:"short_formatted"`
		Formatted      string  `json:"formatted"`
		Address        string  `json:"address"`
		City           string  `json:"city"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Postcode       string  `json:"postcode"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		ZoomLevel      int64   `json:"zoom_level"`
		Approximate    bool    `json:"approximate"`
	} `json:"address"`
	Updated         time.Time `json:"updated"`
	DisableFeedback bool      `json:"disable_feedback"`
}

func bulkTransactionsRequest(auth *Myzo, daysAgo, before int, expandBy, accountId string) *TransactionsResponse {
	r, _ := auth.transactionResponseHandler(true, daysAgo, before, expandBy, "", accountId)
	return r
}

func baseTransactionRequest(auth *Myzo, expandBy, optionalId, accountId string) *TransactionsResponse {
	r, _ := auth.transactionResponseHandler(false, 0, 0, expandBy, optionalId, accountId)
	return r
}

func (auth *Myzo) GetAllTransactions(daysAgo, before int, expandBy, accountId string) []Transaction {
	return bulkTransactionsRequest(auth, daysAgo, before, expandBy, accountId).Transactions
}

func (auth *Myzo) GetTransaction(id, expandBy, accountId string) Transaction {
	return baseTransactionRequest(auth, expandBy, "/"+id, accountId).Transaction
}

func (auth *Myzo) GetAllMerchants(daysAgo, before int, accountId string) []Merchant {
	var merchants []Merchant
	for _, t := range bulkTransactionsRequest(auth, daysAgo, before, "merchant", accountId).Transactions {
		merchants = append(merchants, t.Merchant)
	}
	return merchants
}
