package concurrency

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func MutexCaller() {

	account := NewAccount("Tim")
	if err := account.Deposit(100); err != nil {
		log.Fatal(err)
	}
	account.GetAccountInfo()

	for i := 0; i < 11; i++ {
		go account.Withdraw(10)
	}

	time.Sleep(2 * time.Second)

}

type Account struct {
	name    string
	balance float64
	mutex   sync.Mutex
}

func NewAccount(name string) *Account {
	return &Account{
		name:    name,
		balance: 0,
	}
}

func (a *Account) Withdraw(amount float64) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.balance < amount {
		fmt.Printf("can't be withdrawn, balance less than you try to withdraw\nBalance: %.2f.\nWithdraw Amount: %.2f.\n",
			a.balance, amount)
		return
	}

	a.balance -= amount
	fmt.Printf("Withdrawn: %.2f. Balance left: %.2f\n", amount, a.balance)
}

func (a *Account) Deposit(amount float64) error {
	if amount < 0.1 {
		return fmt.Errorf("wrong amount %.2f, it must be greater that 0.1\n", amount)
	}

	a.balance = amount
	return nil
}

func (a *Account) GetAccountInfo() {
	fmt.Printf("Clinet name: %s.\nClient balance: %.2f$\n", a.name, a.balance)
}
