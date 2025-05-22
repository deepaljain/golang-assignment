package main

import (
	"digital-wallet/wallet"
	"fmt"
)

func main() {
	digitalWallet := digital_wallet.GetDigitalWallet()

	// Create users
	user1 := digital_wallet.NewUser("U001", "John Doe", "john@example.com", "password123")
	user2 := digital_wallet.NewUser("U002", "Jane Smith", "jane@example.com", "password456")
	digitalWallet.CreateUser(user1)
	digitalWallet.CreateUser(user2)

	// Create accounts
	account1 := digital_wallet.NewAccount("A001", user1, "1234567890", digital_wallet.USD)
	account2 := digital_wallet.NewAccount("A002", user2, "9876543210", digital_wallet.USD)
	digitalWallet.CreateAccount(account1)
	digitalWallet.CreateAccount(account2)

	// Deposit funds
	account1.Deposit(1000.00)
	account2.Deposit(500.00)

	// Transfer funds
	amount := 100.00
	if err := digitalWallet.TransferFunds(account1, account2, amount, digital_wallet.USD); err != nil {
		fmt.Printf("Transfer failed: %v\n", err)
	}

	// Get transaction history
	fmt.Println("Transaction History for Account 1:")
	for _, transaction := range digitalWallet.GetTransactionsHistory(account1) {
		fmt.Printf("Transaction ID: %s\n", transaction.ID)
		fmt.Printf("Amount: %v %s\n", transaction.Amount, transaction.Currency)
		fmt.Printf("Transaction Type: %s\n", transaction.Type)
		fmt.Printf("Timestamp: %v\n\n", transaction.Timestamp)
	}

	fmt.Println("Transaction History for Account 2:")
	for _, transaction := range digitalWallet.GetTransactionsHistory(account2) {
		fmt.Printf("Transaction ID: %s\n", transaction.ID)
		fmt.Printf("Amount: %v %s\n", transaction.Amount, transaction.Currency)
		fmt.Printf("Transaction Type: %s\n", transaction.Type)
		fmt.Printf("Timestamp: %v\n\n", transaction.Timestamp)
	}
}