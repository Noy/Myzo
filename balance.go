package myzo

type BalanceResponse struct {
	Balance int64 `json:"balance"`
	TotalBalance int64 `json:"total_balance"`
	BalanceIncludingFlexibleSavings int64 `json:"balance_including_flexible_savings"`
	Currency string `json:"currency"`
	SpendToday int64 `json:"spend_today"`
	LocalCurrency string `json:"local_currency"`
	LocalExchangeRate float64 `json:"local_exchange_rate"`
	LocalSpend []LocalSpend `json:"local_spend"`
}

type LocalSpend struct {
	SpendToday int64 `json:"spend_today"`
	Currency string `json:"currency"`
}

func baseBalanceRequest(auth *Myzo) *BalanceResponse {
	r, _ := auth.balanceResponseHandler()
	return r
}

func (auth *Myzo) BalanceDetails() *BalanceResponse {
	return baseBalanceRequest(auth)
}

func (auth *Myzo) GetBalance() float64 {
	return Convert64IntToFloat(baseBalanceRequest(auth).Balance)
}

func (auth *Myzo) GetTotalBalance() float64 {
	return Convert64IntToFloat(baseBalanceRequest(auth).TotalBalance)
}

func (auth *Myzo) GetSpentToday() float64 {
	return Convert64IntToFloat(baseBalanceRequest(auth).SpendToday)
}

func (auth *Myzo) GetCurrency() string {
	return baseBalanceRequest(auth).Currency
}
