package digital_wallet

import "fmt"

type DigitalWallet struct {}

var users = make(map[string]*User)
var accounts = make(map[string]*Account)

func GetDigitalWallet() *DigitalWallet {
	return &DigitalWallet{}
}

func (dw *DigitalWallet) CreateUser(user *User) {
	if _, exists := users[user.ID]; exists {
		panic("User already exists")
	}
	users[user.ID] = user
}

func (dw *DigitalWallet) CreateAccount(account *Account) {
	if _, exists := accounts[account.ID]; exists {
		panic("Account already exists")
	}
	accounts[account.ID] = account
}

func (dw *DigitalWallet) TransferFunds(from, to *Account, amount float64, currency Currency) error {
	_, exists := accounts[from.ID]
	if !exists {
		return fmt.Errorf("from account not found")
	}

	_, exists = accounts[to.ID]
	if !exists {
		return fmt.Errorf("to account not found")
	}

	if err := NewTransaction(from, to, amount, currency); err != nil {
		return err
	}

	StoreTransaction(from, to, amount, currency)
	fmt.Println("Amount transferred")

	return nil
}

func (dw *DigitalWallet) GetTransactionsHistory(account *Account) []*Transaction {
	transactions := GetTransactionHistory(account)
	if transactions == nil {
		fmt.Println("No transactions found")
		return nil
	}

	return transactions
}