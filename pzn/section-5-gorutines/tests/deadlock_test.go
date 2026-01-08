package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Balance struct {
	Mutex   sync.Mutex
	Nama    string
	Balance int
}

func (user *Balance) Lock() {
	user.Mutex.Lock()
}

func (user *Balance) Unlock() {
	user.Mutex.Unlock()
}

func (user *Balance) ChangeAmount(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *Balance, user2 *Balance, amount int) {
	user1.Lock()
	fmt.Println("Lock ", user1.Nama)
	user1.ChangeAmount(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock ", user2.Nama)
	user2.ChangeAmount(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()

}

func TestDeadlockBalance(t *testing.T) {
	user1 := Balance{
		Nama:    "Budi",
		Balance: 1000000,
	}
	user2 := Balance{
		Nama:    "Rahmat",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(5 * time.Second)

	fmt.Println(user1.Nama, user1.Balance)
	fmt.Println(user2.Nama, user2.Balance)
}
