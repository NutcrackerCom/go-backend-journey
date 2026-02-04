package account

import (
	"errors"
	"sync/atomic"
)

var (
	ErrInvalidAmount     = errors.New("amount must be positive")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNilAccount        = errors.New("account is nil")
)

type Account struct {
	Owner   string
	Balance atomic.Int64
	Log     Logger
}

func (a *Account) logger() Logger {
	if a.Log == nil {
		return NopLogger{}
	}
	return a.Log
}

func (user *Account) Deposit(amount int64) error {
	if amount <= 0 {
		user.logger().Error("Deposit failed, User - %s, amount = %d, Error: %v", user.Owner, amount, ErrInvalidAmount)
		return ErrInvalidAmount
	}
	user.Balance += amount
	user.logger().Info("Deposit ok")
	return nil
}

func (user *Account) Withdraw(amount int64) error {
	if amount <= 0 {
		user.logger().Error("Withdraw failed, User - %s, amount = %d, Error: %v", user.Owner, amount, ErrInvalidAmount)
		return ErrInvalidAmount
	}
	if user.Balance < amount {
		user.logger().Error("Withdraw failed, User - %s, amount = %d, Error: %v", user.Owner, amount, ErrInsufficientFunds)
		return ErrInsufficientFunds
	}
	user.Balance -= amount
	user.logger().Info("Withdraw ok")
	return nil
}

func (user *Account) Transfer(to *Account, amount int64) error {
	if to == nil {
		user.logger().Error("Transfer failed, receiver is nil")
		return ErrNilAccount
	}
	if err := user.Withdraw(amount); err != nil {
		user.logger().Error("Transfer failed, user %s, amount = %d, Error: %v", user.Owner, amount, err)
		return err
	}
	if err := to.Deposit(amount); err != nil {
		_ = user.Deposit(amount)
		user.logger().Error("Transfer failed, user %s, amount = %d, Error: %v", user.Owner, amount, err)
		return err
	}
	user.logger().Info("Transfer ok")
	return nil
}
