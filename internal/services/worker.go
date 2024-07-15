package services

import (
	"fmt"
	"strconv"
	"sync"
)

type Request struct {
	ID     string
	Amount float64
	Result chan error
}

type BalanceRequest struct {
	ID     string
	Result chan float64
}

var accounts = make(map[string]*Account)
var accountsMutex = sync.Mutex{}

var depositChan = make(chan Request)
var withdrawChan = make(chan Request)
var balanceChan = make(chan BalanceRequest)
var createAccChan = make(chan BalanceRequest)

func CreateAccount() string {
	id := strconv.Itoa(len(accounts) + 1)
	account := NewAccount(id)
	accountsMutex.Lock()
	accounts[id] = account
	accountsMutex.Unlock()
	return id
}

func GetAccount(id string) (*Account, error) {
	accountsMutex.Lock()
	defer accountsMutex.Unlock()
	account, exists := accounts[id]
	if !exists {
		return nil, fmt.Errorf("account not found")
	}
	return account, nil
}

func Deposit(id string, amount float64) error {
	result := make(chan error)
	depositChan <- Request{ID: id, Amount: amount, Result: result}
	return <-result
}

func Withdraw(id string, amount float64) error {
	result := make(chan error)
	withdrawChan <- Request{ID: id, Amount: amount, Result: result}
	return <-result
}

func GetBalance(id string) float64 {
	result := make(chan float64)
	balanceChan <- BalanceRequest{ID: id, Result: result}
	return <-result
}

func SetupWorkers(n int) {
	for i := 0; i <= n; i++ {
		go worker()
	}
}

func worker() {
	for {
		select {
		case req := <-depositChan:
			go func(r Request) {
				accountsMutex.Lock()
				account, exists := accounts[r.ID]
				accountsMutex.Unlock()
				if !exists {
					r.Result <- fmt.Errorf("account not found")
					return
				}
				r.Result <- account.Deposit(r.Amount)
			}(req)
		case req := <-withdrawChan:
			go func(r Request) {
				accountsMutex.Lock()
				account, exists := accounts[r.ID]
				accountsMutex.Unlock()
				if !exists {
					r.Result <- fmt.Errorf("account not found")
					return
				}
				r.Result <- account.Withdraw(r.Amount)
			}(req)
		case req := <-balanceChan:
			go func(r BalanceRequest) {
				accountsMutex.Lock()
				account, exists := accounts[r.ID]
				accountsMutex.Unlock()
				if !exists {
					r.Result <- 0
					return
				}
				r.Result <- account.GetBalance()
			}(req)
		}
	}
}
