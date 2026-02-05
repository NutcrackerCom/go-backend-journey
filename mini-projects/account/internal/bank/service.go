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

func (s *Service) Create(id int, owner string, initialBalance int64) error {
	if initialBalance < 0 {
		s.log.Error("Eror: initialBalance less then zero")
		return ErrInvalidAmount
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.account[id]; ok {
		s.log.Error("Eror: Try to create existing account")
		return ErrAlreadyExists
	}
	if _, ok := s.locks[id]; ok {
		s.log.Error("Eror: Try to create existing locks")
		return ErrAlreadyExistsLock
	}
	user := model.Account{Id: id, Owner: owner, Balance: initialBalance}
	s.account[id] = &user
	s.locks[id] = &sync.Mutex{}
	return nil
}

func (s *Service) Get(id int) (model.Account, error) {
	s.mu.RLock()
	user, ok := s.account[id]
	lock := s.locks[id]
	s.mu.RUnlock()
	if !ok || lock == nil {
		return model.Account{}, ErrNotFound
	}
	lock.Lock()
	acc := *user
	lock.Unlock()
	return acc, nil
}

func (s *Service) getLockAndAcc(accId int) (*sync.Mutex, *model.Account, error) {
	s.mu.RLock()
	acc, ok_acc := s.account[accId]
	lock, ok_lock := s.locks[accId]
	s.mu.RUnlock()
	if acc == nil || !ok_acc {
		return nil, nil, ErrNotFound
	}
	if lock == nil || !ok_lock {
		return nil, nil, ErrNotFound
	}
	return lock, acc, nil
}

func (s *Service) deposit(id int, amount int64) error {
	lock, user, err := s.getLockAndAcc(id)
	if err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	return s.depositLocked(user, amount)
}

func (s *Service) depositLocked(user *model.Account, amount int64) error {
	if amount <= 0 {
		s.log.Error("Deposit failed")
		return ErrInvalidAmount
	}

	user.Balance += amount
	s.log.Info("Deposit ok")
	return nil
}

func (s *Service) withdraw(id int, amount int64) error {
	lock, user, err := s.getLockAndAcc(id)
	if err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	return s.withdrawLocked(user, amount)
}

func (s *Service) withdrawLocked(user *model.Account, amount int64) error {
	if amount <= 0 {
		s.log.Error("Withdraw failed")
		return ErrInvalidAmount
	}

	if user.Balance < amount {
		s.log.Error("Withdraw failed")
		return ErrInsufficientFunds
	}

	user.Balance -= amount
	s.log.Info("Withdraw ok")
	return nil
}

func (s *Service) Transfer(fromId int, toId int, amount int64) error {
	if fromId == toId {
		s.log.Error("Eror: duplicate from id = %d to id = %d ", fromId, toId)
		return ErrDuplicateId
	}

	fromLock, userFrom, err := s.getLockAndAcc(fromId)
	if err != nil {
		return err
	}

	toLock, userTo, err := s.getLockAndAcc(toId)
	if err != nil {
		return err
	}

	firstLock, secondLock := fromLock, toLock
	if fromId > toId {
		secondLock, firstLock = fromLock, toLock
	}

	firstLock.Lock()
	defer firstLock.Unlock()
	secondLock.Lock()
	defer secondLock.Unlock()

	if err := s.withdrawLocked(userFrom, amount); err != nil {
		s.log.Error("Transfer failed")
		return err
	}
	if err := s.depositLocked(userTo, amount); err != nil {
		s.depositLocked(userFrom, amount)
		s.log.Error("Transfer failed")
		return err
	}
	s.log.Info("Transfer ok")
	return nil
}
