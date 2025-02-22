package sixth

import (
	"errors"
	"fmt"
)

var ErrorInsufficientFunds = errors.New("insufficient funds")

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrorInsufficientFunds
	}

	w.balance -= amount
	return nil
}
