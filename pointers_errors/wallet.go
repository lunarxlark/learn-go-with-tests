package pointers_errors

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(money Bitcoin) {
	w.balance += money
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
