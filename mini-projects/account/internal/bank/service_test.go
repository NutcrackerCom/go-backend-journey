package bank

import (
	"errors"
	"testing"
)

func TestDeposit_Positive(t *testing.T) {
	accounts := NewService(nil)
	accounts.Create(0, "A", 100)
	err := accounts.Deposit(0, 50)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if balance, _ := accounts.Get(0); balance.Balance != 150 {
		t.Fatalf("balance: got %d want %d", balance.Balance, 150)
	}
}

func TestDeposit_NegativeAmount(t *testing.T) {

	accounts := NewService(nil)
	accounts.Create(0, "A", 100)

	err := accounts.Deposit(0, -10)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if balance, _ := accounts.Get(0); balance.Balance != 100 {
		t.Fatalf("balance must not change on error: got %d want %d", balance.Balance, 100)
	}
}

func TestWithdraw_EnoughFunds(t *testing.T) {
	accounts := NewService(nil)
	accounts.Create(0, "A", 100)
	err := accounts.Withdraw(0, 60)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if balance, _ := accounts.Get(0); balance.Balance != 40 {
		t.Fatalf("balance: got %d want %d", balance.Balance, 40)
	}
}

func TestWithdraw_NotEnoughFunds(t *testing.T) {
	accounts := NewService(nil)
	accounts.Create(0, "A", 50)

	err := accounts.Withdraw(0, 60)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if balance, _ := accounts.Get(0); balance.Balance != 50 {
		t.Fatalf("balance must not change on error: got %d want %d", balance.Balance, 50)
	}
}

func TestTransfer_Success(t *testing.T) {
	accounts := NewService(nil)
	accounts.Create(0, "A", 100)
	accounts.Create(1, "B", 10)
	err := accounts.Transfer(0, 1, 70)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if balance, _ := accounts.Get(0); balance.Balance != 30 {
		t.Fatalf("from balance: got %d want %d", balance.Balance, 30)
	}
	if balance, _ := accounts.Get(1); balance.Balance != 80 {
		t.Fatalf("to balance: got %d want %d", balance.Balance, 80)
	}
}

func TestTransfer_NotEnoughFunds(t *testing.T) {
	accounts := NewService(nil)
	accounts.Create(0, "A", 20)
	accounts.Create(1, "B", 10)

	err := accounts.Transfer(0, 1, 30)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if balance, _ := accounts.Get(0); balance.Balance != 20 {
		t.Fatalf("from balance must not change on error: got %d want %d", balance.Balance, 20)
	}
	if balance, _ := accounts.Get(1); balance.Balance != 10 {
		t.Fatalf("to balance must not change on error: got %d want %d", balance.Balance, 10)
	}
}

func TestTransfer_NilAccount(t *testing.T) {
	accounts := NewService(nil)
	accounts.Create(0, "A", 20)
	err := accounts.Transfer(0, 1, 30)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected error ErrUnExsisingUser, got %v", err)
	}
}

func TestDeposit_Zero(t *testing.T) {
	accounts := NewService(nil)
	accounts.Create(0, "A", 20)
	err := accounts.Deposit(0, 0)
	if !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected error ErrInvalidAmount, got %v", err)
	}
}

func TestWithdraw_Zero(t *testing.T) {
	accounts := NewService(nil)
	accounts.Create(0, "A", 20)
	err := accounts.Withdraw(0, 0)
	if !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected error ErrInvalidAmount, got %v", err)
	}
}
