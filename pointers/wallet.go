package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin is cool currency
type Bitcoin int

// Wallet holds moneys
type Wallet struct {
	balance Bitcoin
}

// ErrInsufficientFunds is thrown when there are not enough funds in wallet
var ErrInsufficientFunds = errors.New("Insufficient funds")

// Deposit puts money in a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns how much money is in a wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Withdraw takes money out of a wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
