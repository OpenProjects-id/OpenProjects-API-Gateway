package transaction

import "time"

type ProjectTransactionsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatProjectTransaction(transaction Transaction) ProjectTransactionsFormatter {
	formatter := ProjectTransactionsFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatProjectTransactions(transactions []Transaction) []ProjectTransactionsFormatter {
	if len(transactions) == 0 {
		return []ProjectTransactionsFormatter{}
	}

	var transactionsFormatter []ProjectTransactionsFormatter

	for _, transaction := range transactions {
		formatter := FormatProjectTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}
