package todo

import (
	"errors"
	"sort"
	"sync"
	"time"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrEmptyText = errors.New("empty text")
)

type Service struct {
	nextID int
	tasks  map[int]Task
	mu     sync.Mutex
}

func NewService() *Service {
	return &Service{
		nextID: 0,
		tasks:  map[int]Task{},
		mu:     sync.Mutex{}}
}

func (s *Service) Add(title string) (Task, error) {
	if title == "" {
		return Task{}, ErrEmptyText
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	task := Task{
		ID:        s.nextID,
		Text:      title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	s.tasks[s.nextID] = task
	s.nextID++
	return task, nil
}

func (s *Service) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	var sliceTasks []Task
	for _, task := range s.tasks {
		sliceTasks = append(sliceTasks, task)
	}
	sort.Slice(sliceTasks, func(i, j int) bool {
		return sliceTasks[i].ID < sliceTasks[j].ID
	})
	return sliceTasks
}

func (s *Service) Done(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tasks[id]; !ok {
		return ErrNotFound
	}
	t := s.tasks[id]
	t.Done = true
	s.tasks[id] = t
	return nil
}

func (s *Service) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tasks[id]; !ok {
		return ErrNotFound
	}
	delete(s.tasks, id)
	return nil
}
