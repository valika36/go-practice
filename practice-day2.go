package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {

	b.mu.Lock()
	defer b.mu.Unlock()

	fmt.Println("current balance: ", b.balance)
	b.balance += amount
	fmt.Printf("deposit: %d, current balance: %d\n", amount, b.balance)
}

func (b *BankAccount) WithDraw(amount int) {

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.balance < amount {
		fmt.Println("Insufficient balance: ", b.balance)
		return
	}

	fmt.Println("current balance: ", b.balance)
	b.balance -= amount
	fmt.Printf("withdraw: %d, current balance: %d\n", amount, b.balance)
}

func main() {

	account := &BankAccount{balance: 500}

	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(1 * time.Second)
			account.Deposit(300)
		}
	}()
	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(1 * time.Second)
			account.WithDraw(100)
		}
	}()
	time.Sleep(time.Second * 10)
	fmt.Println("Final balance:", account.balance)
}
