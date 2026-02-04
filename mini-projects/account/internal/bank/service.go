package bank

import (
	"errors"
	"sync"

	"github.com/NutcrackerCom/go-backend-journey/mini-projects/account/internal/logger"
	"github.com/NutcrackerCom/go-backend-journey/mini-projects/account/internal/model"
)

var (
	ErrInvalidAmount     = errors.New("amount must be positive")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNotFound          = errors.New("user do not exists")
	ErrAlreadyExists     = errors.New("create existing account")
	ErrAlreadyExistsLock = errors.New("create existing lock")
	ErrDuplicateId       = errors.New("duplicate ids")
)

func NewService(log logger.Logger) *Service {
	if log == nil {
		log = logger.NopLogger{}
	}

	return &Service{
		mu:      sync.RWMutex{},
		account: make(map[int]*model.Account),
		locks:   make(map[int]*sync.Mutex),
		log:     log,
	}
}

type Service struct {
	mu sync.RWMutex

	account map[int]*model.Account
	locks   map[int]*sync.Mutex
	log     logger.Logger
}

func (accounts *Service) Create(id int, owner string, initialBalance int64) error {

	if _, ok := accounts.account[id]; ok {
		accounts.log.Error("Eror: Try to create existing account")
		return ErrAlreadyExists
	}
	if _, ok := accounts.locks[id]; ok {
		accounts.log.Error("Eror: Try to create existing locks")
		return ErrAlreadyExistsLock
	}
	if initialBalance < 0 {
		accounts.log.Error("Eror: initialBalance less then zero")
		return ErrInvalidAmount
	}
	user := model.Account{Id: id, Owner: owner, Balance: initialBalance}
	accounts.account[id] = &user
	accounts.locks[id] = &sync.Mutex{}
	return nil
}

func (accounts *Service) Get(id int) (model.Account, error) {
	user, ok := accounts.account[id]
	if ok {
		return *user, nil
	} else {
		return model.Account{}, ErrNotFound
	}
}

func (accounts *Service) Deposit(id int, amount int64) error {
	if _, ok := accounts.account[id]; !ok {
		accounts.log.Error("Eror: user do not exists")
		return ErrNotFound
	}
	if amount <= 0 {
		accounts.log.Error("Deposit failed")
		return ErrInvalidAmount
	}
	accounts.account[id].Balance += amount
	return nil
}

func (accounts *Service) Withdraw(id int, amount int64) error {
	user, ok := accounts.account[id]
	if !ok {
		accounts.log.Error("Eror: user do not exists")
		return ErrNotFound
	}
	if amount <= 0 {
		accounts.log.Error("Withdraw failed")
		return ErrInvalidAmount
	}
	if user.Balance < amount {
		accounts.log.Error("Withdraw failed")
		return ErrInsufficientFunds
	}
	user.Balance -= amount
	accounts.log.Info("Withdraw ok")
	return nil
}

func (accounts *Service) Transfer(from_id int, to_id int, amount int64) error {
	if from_id == to_id {
		accounts.log.Error("Eror: duplicate from id = %d to id = %d ", from_id, to_id)
		return ErrDuplicateId
	}
	if _, ok := accounts.account[from_id]; !ok {
		accounts.log.Error("Eror: user with id = %d do not exists", from_id)
		return ErrNotFound
	}
	if _, ok := accounts.account[to_id]; !ok {
		accounts.log.Error("Eror: user with id = %d do not exists", to_id)
		return ErrNotFound
	}
	if err := accounts.Withdraw(from_id, amount); err != nil {
		accounts.log.Error("Transfer failed")
		return err
	}
	if err := accounts.Deposit(to_id, amount); err != nil {
		accounts.Deposit(from_id, amount)
		accounts.log.Error("Transfer failed")
		return err
	}
	accounts.log.Info("Transfer ok")
	return nil
}
