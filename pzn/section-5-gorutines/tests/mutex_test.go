package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceMutexCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}

		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Count : ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println("Get Balance : ", account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance : ", account.GetBalance())
}
