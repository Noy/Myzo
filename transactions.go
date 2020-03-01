package myzo

func bulkTransactionsRequest(auth *Myzo, daysAgo int, expandBy, limit string) *TransactionsResponse {
	r, _ := auth.transactionResponseHandler(true, daysAgo, expandBy, "", limit)
	return r
}

func baseTransactionRequest(auth *Myzo, expandBy, optionalId string) *TransactionsResponse {
	r, _ := auth.transactionResponseHandler(false,0, expandBy, optionalId, "")
	return r
}

func (auth *Myzo) GetAllTransactions(daysAgo int, expandBy, limit string) []Transaction {
	return bulkTransactionsRequest(auth, daysAgo, expandBy, limit).Transactions
}

func (auth *Myzo) GetTransaction(id, expandBy string) Transaction {
	return baseTransactionRequest(auth, expandBy, "/"+id).Transaction
}

func (auth *Myzo) GetAllMerchants(daysAgo int, limit string) []Merchant {
	var merchants []Merchant
	for _, t := range bulkTransactionsRequest(auth, daysAgo, "merchant", limit).Transactions {
		merchants = append(merchants, t.Merchant)
	}
	return merchants
}