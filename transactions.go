package myzo

func bulkTransactionsRequest(auth *Myzo, daysAgo int, expandBy string) *TransactionsResponse {
	r, _ := auth.transactionResponseHandler(true, daysAgo, expandBy, "")
	return r
}

func baseTransactionRequest(auth *Myzo, expandBy, optionalId string) *TransactionsResponse {
	r, _ := auth.transactionResponseHandler(false,0, expandBy, optionalId)
	return r
}

func (auth *Myzo) GetAllTransactions(daysAgo int, expandBy string) []Transaction {
	return bulkTransactionsRequest(auth, daysAgo, expandBy).Transactions
}

func (auth *Myzo) GetTransaction(id, expandBy string) Transaction {
	return baseTransactionRequest(auth, expandBy, "/"+id).Transaction
}

func (auth *Myzo) GetAllMerchants(daysAgo int, expandBy string) *Merchant {
	for _, t := range bulkTransactionsRequest(auth, daysAgo, expandBy).Transactions {
		return &t.Merchant
	}
	return nil
}
