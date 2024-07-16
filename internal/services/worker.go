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

var (
	accounts      = make(map[string]*Account)
	accountsMutex sync.Mutex
)

var depositChan = make(chan Request)
var withdrawChan = make(chan Request)
var balanceChan = make(chan BalanceRequest)

func CreateAccount() string {
	id := strconv.Itoa(len(accounts) + 1)
	account := NewAccount(id)
	accountsMutex.Lock()
	defer accountsMutex.Unlock()
	accounts[id] = account
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
				acc, err := GetAccount(r.ID)
				if err != nil {
					r.Result <- err
					return
				}
				r.Result <- acc.Deposit(r.Amount)
			}(req)
		case req := <-withdrawChan:
			go func(r Request) {
				acc, err := GetAccount(r.ID)
				if err != nil {
					r.Result <- err
					return
				}
				r.Result <- acc.Withdraw(r.Amount)
			}(req)
		case req := <-balanceChan:
			go func(r BalanceRequest) {
				acc, err := GetAccount(r.ID)
				if err != nil {
					r.Result <- 0
					return
				}
				r.Result <- acc.GetBalance()
			}(req)
		}
	}
}
