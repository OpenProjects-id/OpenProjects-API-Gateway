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

type UserTransactionFormatter struct {
	ID        int              `json:"id"`
	Amount    int              `json:"amount"`
	Status    string           `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	Project   ProjectFormatter `json:"project"`
}

type ProjectFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	projectFormatter := ProjectFormatter{}
	projectFormatter.Name = transaction.Project.Name
	projectFormatter.ImageURL = ""

	if len(transaction.Project.ProjectImages) > 0 {
		projectFormatter.ImageURL = transaction.Project.ProjectImages[0].FileName
	}

	formatter.Project = projectFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}
