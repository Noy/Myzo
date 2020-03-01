package myzo

func baseBalanceRequest(auth *Myzo) *BalanceResponse {
	r, _ := auth.balanceResponseHandler()
	return r
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
