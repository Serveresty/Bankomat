package services

import (
	"fmt"
	"log"
	"time"
)

func NewAccount(id string) *Account {
	return &Account{ID: id, Balance: 0.0}
}

func (a *Account) Deposit(amount float64) error {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	a.Balance += amount
	log.Printf("Account %s: Deposited %f at %v\n", a.ID, amount, time.Now())
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	if amount > a.Balance {
		return fmt.Errorf("not enough money")
	}
	a.Balance -= amount
	log.Printf("Account %s: Withdrawn %f at %v\n", a.ID, amount, time.Now())
	return nil
}

func (a *Account) GetBalance() float64 {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	log.Printf("Account %s: Got balance at %v\n", a.ID, time.Now())
	return a.Balance
}
