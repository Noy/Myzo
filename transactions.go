package monzo

func bulkTransactionsRequest(auth *AuthMonzo, daysAgo int, expandBy string) *TransactionsResponse {
	r, _ := auth.transactionResponseHandler(true, daysAgo, expandBy, "")
	return r
}

func baseTransactionRequest(auth *AuthMonzo, expandBy, optionalId string) *TransactionsResponse {
	r, _ := auth.transactionResponseHandler(false,0, expandBy, optionalId)
	return r
}

func (auth *AuthMonzo) GetAllTransactions(daysAgo int, expandBy string) []Transaction {
	return bulkTransactionsRequest(auth, daysAgo, expandBy).Transactions
}

func (auth *AuthMonzo) GetTransaction(id, expandBy string) Transaction {
	return baseTransactionRequest(auth, "merchant", "/"+id).Transaction
}

func (auth *AuthMonzo) GetAllMerchants(daysAgo int, expandBy string) *Merchant {
	for _, t := range bulkTransactionsRequest(auth, daysAgo, expandBy).Transactions {
		return &t.Merchant
	}
	return nil
}
