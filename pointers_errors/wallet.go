package pointers_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(money Bitcoin) {
	w.balance += money
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(money Bitcoin) error {
	if money > w.balance {
		return errors.New("wallet less than money")
	}
	w.balance -= money
	return nil
}
