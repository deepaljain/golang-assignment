package digital_wallet

import "fmt"

type Account struct {
	ID					string
	User				*User
	AccountNumber		string
	Balance				float64
	Currency			Currency
}

func NewAccount(id string, user *User, accountNumber string, currency Currency) *Account {
	return &Account{
		ID:				id,
		User:			user,
		AccountNumber:	accountNumber,
		Balance:		0.0,
		Currency:		currency,
	}
}

func (account *Account) Deposit(amount float64) {
	account.Balance += amount
}

func (account *Account) Withdraw(amount float64) error {
	if amount > account.Balance {
		return fmt.Errorf("insufficient funds")
	}
	account.Balance -= amount
	return nil
}