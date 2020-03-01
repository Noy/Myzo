package myzo

func baseAccountRequestRequest(auth *Myzo) *AccountResponse {
	r, _ := auth.accountResponseHandler()
	return r
}

func (auth *Myzo) GetAccounts() []Account {
	return baseAccountRequestRequest(auth).Accounts
}