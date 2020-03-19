package myzo

type BalanceResponse struct {
	Balance                         int64        `json:"balance"`
	TotalBalance                    int64        `json:"total_balance"`
	BalanceIncludingFlexibleSavings int64        `json:"balance_including_flexible_savings"`
	Currency                        string       `json:"currency"`
	SpendToday                      int64        `json:"spend_today"`
	LocalCurrency                   string       `json:"local_currency"`
	LocalExchangeRate               float64      `json:"local_exchange_rate"`
	LocalSpend                      []LocalSpend `json:"local_spend"`
}

type LocalSpend struct {
	SpendToday int64  `json:"spend_today"`
	Currency   string `json:"currency"`
}

func baseBalanceRequest(auth *Myzo, accountId string) *BalanceResponse {
	r, _ := auth.balanceResponseHandler(accountId)
	return r
}

func (auth *Myzo) BalanceDetails(accountId string) *BalanceResponse {
	return baseBalanceRequest(auth, accountId)
}

func (auth *Myzo) GetBalance(accountId string) float64 {
	return Convert64IntToFloat(baseBalanceRequest(auth, accountId).Balance)
}

func (auth *Myzo) GetTotalBalance(accountId string) float64 {
	return Convert64IntToFloat(baseBalanceRequest(auth, accountId).TotalBalance)
}

func (auth *Myzo) GetSpentToday(accountId string) float64 {
	return Convert64IntToFloat(baseBalanceRequest(auth, accountId).SpendToday)
}

func (auth *Myzo) GetCurrency(accountId string) string {
	return baseBalanceRequest(auth, accountId).Currency
}
