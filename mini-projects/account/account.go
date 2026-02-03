package account

import (
	"errors"
)

var (
	ErrInvalidAmount     = errors.New("amount must be positive")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNilAccount        = errors.New("Account is nil")
)

type Account struct {
	Owner   string
	Balance int64
	log     Logger
}

func (user *Account) Deposit(amount int64) error {
	if amount <= 0 {
		user.log.Error("Deposit failed")
		return ErrInvalidAmount
	}
	user.Balance += amount
	user.log.Info("Deposit ok")
	return nil
}

func (user *Account) Withdraw(amount int64) error {
	if amount <= 0 {
		user.log.Error("Withdraw failed")
		return ErrInvalidAmount
	}
	if user.Balance < amount {
		user.log.Error("Withdraw failed")
		return ErrInsufficientFunds
	}
	user.Balance -= amount
	user.log.Info("Withdraw ok")
	return nil
}

func (user *Account) Transfer(to *Account, amount int64) error {
	if to == nil {
		user.log.Error("Transfer failed")
		return ErrNilAccount
	}
	if err := user.Withdraw(amount); err != nil {
		user.log.Error("Transfer failed")
		return err
	}
	if err := to.Deposit(amount); err != nil {
		_ = user.Deposit(amount)
		user.log.Error("Transfer failed")
		return err
	}
	user.log.Info("Transfer ok")
	return nil
}
