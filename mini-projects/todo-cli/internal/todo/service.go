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
	mu     *sync.Mutex
}

func NewService() *Service {
	return &Service{
		nextID: 0,
		tasks:  map[int]Task{},
		mu:     &sync.Mutex{}}
}

func (s *Service) Add(title string) (Task, error) {
	if title == "" {
		return Task{}, ErrEmptyText
	}
	task := Task{
		ID:        s.nextID,
		Text:      title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	s.mu.Lock()
	s.tasks[s.nextID] = task
	s.nextID++
	s.mu.Unlock()
	return task, nil
}

func (s *Service) List() []Task {
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
	if _, ok := s.tasks[id]; !ok {
		return ErrNotFound
	}
	s.tasks[id] = Task{
		ID:        s.tasks[id].ID,
		Text:      s.tasks[id].Text,
		Done:      true,
		CreatedAt: s.tasks[id].CreatedAt,
	}
	s.mu.Unlock()
	return nil
}

func (s *Service) Delete(id int) error {
	s.mu.Lock()
	if _, ok := s.tasks[id]; !ok {
		return ErrNotFound
	}
	delete(s.tasks, id)
	s.mu.Unlock()
	return nil
}
