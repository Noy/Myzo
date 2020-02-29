package myzo

func baseBalanceRequest(auth *AuthMonzo) *BalanceResponse {
	r, _ := auth.balanceResponseHandler()
	return r
}

func (auth *AuthMonzo) GetBalance() float64 {
	return Convert64IntToFloat(baseBalanceRequest(auth).Balance)
}

func (auth *AuthMonzo) GetSpentToday() float64 {
	return Convert64IntToFloat(baseBalanceRequest(auth).SpendToday)
}
