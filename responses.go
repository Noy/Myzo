package monzo

import "time"

type BalanceResponse struct {
	Balance int64 `json:"balance"`
	TotalBalance int64 `json:"total_balance"`
	Currency string `json:"currency"`
	SpendToday int64 `json:"spend_today"`
}

type PotResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Style string `json:"style"`
	Balance int64 `json:"balance"`
	Currency string `json:"currency"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Deleted bool `json:"deleted"`
}

type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
	Transaction Transaction `json:"transaction"`
}

type Transaction struct {
	AccountBalance int64 `json:"account_balance"`
	Amount int64 `json:"amount"`
	Created time.Time `json:"created"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	ID string `json:"id"`
	MetaData struct{} `json:"metadata"`
	Notes string `json:"notes"`
	IsLoad bool `json:"is_load"`
	Settled string `json:"settled"`
	Category string `json:"category"`
	LocalAmount int64 `json:"local_amount"`
	LocalCurrency string `json:"local_currency"`
	CanBeExcludedFromBreakdown bool `json:"can_be_excluded_from_breakdown"`
	CanBeMadeSubscription bool `json:"can_be_made_subscription"`
	CanSplitBill bool `json:"can_split_the_bill"`
	CanAddToTab bool `json:"can_add_to_tab"`
	AmountIsPending bool `json:"amount_is_pending"`
	Originator bool `json:"originator"`
	DedupeID string `json:"dedupe_id"`
	Scheme string `json:"scheme"`
	Merchant Merchant `json:"merchant"`
}

type Merchant struct {
	ID string `json:"id"`
	GroupID string `json:"group_id"`
	Created time.Time `json:"created"`
	Name string `json:"name"`
	Logo string `json:"logo"`
	Emoji string `json:"emoji"`
	Category string `json:"category"`
	Online bool `json:"online"`
	ATM bool `json:"atm"`
	Address struct {
		ShortFormatted string `json:"short_formatted"`
		Formatted string `json:"formatted"`
		Address string `json:"address"`
		City string `json:"city"`
		Region string `json:"region"`
		Country string `json:"country"`
		Postcode string `json:"postcode"`
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		ZoomLevel int64 `json:"zoom_level"`
		Approximate bool `json:"approximate"`
	} `json:"address"`
	Updated time.Time `json:"updated"`
	DisableFeedback bool `json:"disable_feedback"`
}