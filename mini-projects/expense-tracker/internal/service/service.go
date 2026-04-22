package service

import (
	"errors"
	"sync"
	"time"
)

type ExpenseType string

type ExpensesService struct {
	mu           sync.RWMutex
	ExpensesList map[ExpenseType][]Expenses
	nextId       int
}

const (
	Meal    ExpenseType = "MEAL"
	Clothes ExpenseType = "CLOTHES"
	Study   ExpenseType = "STUDY"
)

var (
	ErrorEmptyTitle        = errors.New("empty title")
	ErrorIncorrectCategory = errors.New("incorrect category")
	ErrorIncorrectAmount   = errors.New("incorrect amount")
	ErrorIdNotFound        = errors.New("id not found")
)

type Expenses struct {
	Id        int         `json:"id"`
	Title     string      `json:"title"`
	Amount    int         `json:"amount"`
	Category  ExpenseType `json:"category"`
	CreatedAt time.Time   `json:"created_at"`
}

func (s *ExpensesService) List() map[ExpenseType][]Expenses {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make(map[ExpenseType][]Expenses, len(s.ExpensesList))
	for k, v := range s.ExpensesList {
		out[k] = append([]Expenses(nil), v...)
	}

	return out
}

func (s *ExpensesService) Add(expenses Expenses) error {
	if expenses.Title == "" {
		return ErrorEmptyTitle
	}
	if expenses.Category == "" || (expenses.Category != Meal &&
		expenses.Category != Clothes && expenses.Category != Study) {
		return ErrorIncorrectCategory
	}
	if expenses.Amount <= 0 {
		return ErrorIncorrectAmount
	}

	expenses.CreatedAt = time.Now()

	s.mu.Lock()
	defer s.mu.Unlock()

	expenses.Id = s.nextId
	s.nextId++
	s.ExpensesList[expenses.Category] = append(s.ExpensesList[expenses.Category], expenses)
	return nil
}

func (s *ExpensesService) DeleteById(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for typeExp, listExp := range s.ExpensesList {
		for ind, exp := range listExp {
			if exp.Id == id {
				s.ExpensesList[typeExp] = append(s.ExpensesList[typeExp][:ind], s.ExpensesList[typeExp][ind+1:]...)
				return nil
			}
		}
	}
	return ErrorIdNotFound
}
