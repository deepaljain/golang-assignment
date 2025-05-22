package digital_wallet

import "digital-wallet/utils"

type TransactionType string

const (
	DEBIT TransactionType = "debit"
	CREDIT TransactionType = "credit"
)

type Transaction struct {
	ID          string
	Account     *Account
	Amount      float64
	Currency	Currency
	Type        TransactionType
	Timestamp   string
}

var transactions = make(map[string][]*Transaction)

func NewTransaction(from *Account, to *Account, amount float64, currency Currency) error {
	if err := from.Withdraw(amount); err != nil {
		return err
	}

	to.Deposit(amount)
	return nil
}

func StoreTransaction(from *Account, to *Account, amount float64, currency Currency) {
	creditTransaction := &Transaction{
		ID:        utils.GenerateTransactionID(),
		Account:   to,
		Amount:    amount,
		Currency:  currency,
		Type:      CREDIT,
		Timestamp: utils.GetCurrentTimestamp(),
	}

	debitTransaction := &Transaction{
		ID:        utils.GenerateTransactionID(),
		Account:   from,
		Amount:    amount,
		Currency:  currency,
		Type:      DEBIT,
		Timestamp: utils.GetCurrentTimestamp(),
	}

	transactions[from.ID] = append(transactions[from.ID], debitTransaction)
	transactions[to.ID] = append(transactions[to.ID], creditTransaction)
}

func GetTransactionHistory(account *Account) []*Transaction {
	accountTransactions, exists := transactions[account.ID]
	if !exists {
		return nil
	}
	return accountTransactions
}