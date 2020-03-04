package myzo

func bulkTransactionsRequest(auth *Myzo, daysAgo,before int, expandBy, limit string) *TransactionsResponse {
	r, _ := auth.transactionResponseHandler(true, daysAgo, before, expandBy, "", limit)
	return r
}

func baseTransactionRequest(auth *Myzo, expandBy, optionalId string) *TransactionsResponse {
	r, _ := auth.transactionResponseHandler(false,0,0, expandBy, optionalId, "")
	return r
}

func (auth *Myzo) GetAllTransactions(daysAgo, before int, expandBy, limit string) []Transaction {
	return bulkTransactionsRequest(auth, daysAgo, before, expandBy, limit).Transactions
}

func (auth *Myzo) GetTransaction(id, expandBy string) Transaction {
	return baseTransactionRequest(auth, expandBy, "/"+id).Transaction
}

func (auth *Myzo) GetAllMerchants(daysAgo, before int, limit string) []Merchant {
	var merchants []Merchant
	for _, t := range bulkTransactionsRequest(auth, daysAgo, before, "merchant", limit).Transactions {
		merchants = append(merchants, t.Merchant)
	}
	return merchants
}