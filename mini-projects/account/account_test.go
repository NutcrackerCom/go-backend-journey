package account

import (
	"errors"
	"testing"
)

func TestDeposit_Positive(t *testing.T) {
	acc := Account{Owner: "A", Balance: 100}

	err := acc.Deposit(50)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if acc.Balance != 150 {
		t.Fatalf("balance: got %d want %d", acc.Balance, 150)
	}
}

func TestDeposit_NegativeAmount(t *testing.T) {
	acc := Account{Owner: "A", Balance: 100}

	err := acc.Deposit(-10)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if acc.Balance != 100 {
		t.Fatalf("balance must not change on error: got %d want %d", acc.Balance, 100)
	}
}

func TestWithdraw_EnoughFunds(t *testing.T) {
	acc := Account{Owner: "A", Balance: 100}

	err := acc.Withdraw(60)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if acc.Balance != 40 {
		t.Fatalf("balance: got %d want %d", acc.Balance, 40)
	}
}

func TestWithdraw_NotEnoughFunds(t *testing.T) {
	acc := Account{Owner: "A", Balance: 50}

	err := acc.Withdraw(60)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if acc.Balance != 50 {
		t.Fatalf("balance must not change on error: got %d want %d", acc.Balance, 50)
	}
}

func TestTransfer_Success(t *testing.T) {
	from := Account{Owner: "A", Balance: 100}
	to := Account{Owner: "B", Balance: 10}

	err := from.Transfer(&to, 70)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if from.Balance != 30 {
		t.Fatalf("from balance: got %d want %d", from.Balance, 30)
	}
	if to.Balance != 80 {
		t.Fatalf("to balance: got %d want %d", to.Balance, 80)
	}
}

func TestTransfer_NotEnoughFunds(t *testing.T) {
	from := Account{Owner: "A", Balance: 20}
	to := Account{Owner: "B", Balance: 10}

	err := from.Transfer(&to, 30)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if from.Balance != 20 {
		t.Fatalf("from balance must not change on error: got %d want %d", from.Balance, 20)
	}
	if to.Balance != 10 {
		t.Fatalf("to balance must not change on error: got %d want %d", to.Balance, 10)
	}
}

func TestTransfer_NilAccount(t *testing.T) {
	from := Account{Owner: "A", Balance: 20}
	err := from.Transfer(nil, 30)
	if !errors.Is(err, ErrNilAccount) {
		t.Fatalf("expected error ErrNilAccount, got %v", err)
	}
}

func TestDeposit_Zero(t *testing.T) {
	from := Account{Owner: "A", Balance: 20}
	err := from.Deposit(0)
	if !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected error ErrInvalidAmount, got %v", err)
	}
}

func TestWithdraw_Zero(t *testing.T) {
	from := Account{Owner: "A", Balance: 20}
	err := from.Withdraw(0)
	if !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected error ErrInvalidAmount, got %v", err)
	}
}
