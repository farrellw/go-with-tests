package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

const minimumBalance Bitcoin = Bitcoin(0)

var ErrInsufficientFunds = errors.New("Can not withdraw past minimum balance")

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if (w.balance - amount) < minimumBalance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
