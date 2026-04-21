package service

import (
	"errors"
	"time"
)

type TaskStatus string

type TaskService struct {
	TaskList []Task
	nextId   int
}

const (
	Done    TaskStatus = "DONE"
	Todo    TaskStatus = "TODO"
	Suspend TaskStatus = "SUSPEND"
)

var (
	ErrorEmptyTaskName   = errors.New("empty task name")
	ErrorIncorrectStatus = errors.New("incorrect task status")
	ErrorIdNotFound      = errors.New("id not found")
)

type Task struct {
	Id        int        `json:"id"`
	Task      string     `json:"task"`
	CreatedAt time.Time  `json:"created_at"`
	Status    TaskStatus `json:"status"`
}

func (s *TaskService) List() []Task {
	return s.TaskList
}

func (s *TaskService) Add(task Task) error {
	if task.Task == "" {
		return ErrorEmptyTaskName
	}
	if task.Status == "" || (task.Status != Done && task.Status != Todo && task.Status != Suspend) {
		return ErrorIncorrectStatus
	}
	task.CreatedAt = time.Now()
	task.Id = s.nextId
	s.nextId++
	s.TaskList = append(s.TaskList, task)
	return nil
}

func (s *TaskService) DeleteById(id int) error {
	for ind, task := range s.TaskList {
		if task.Id == id {
			s.TaskList = append(s.TaskList[:ind], s.TaskList[ind+1:]...)
			return nil
		}
	}
	return ErrorIdNotFound
}
