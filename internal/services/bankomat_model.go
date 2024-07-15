package services

import "sync"

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	ID      string
	Balance float64
	Mutex   sync.Mutex
}
