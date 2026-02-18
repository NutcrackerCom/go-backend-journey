package todo

import (
	"errors"
	"testing"
)

func compare(ref Service, comp Service) bool {
	if ref.nextID != comp.nextID {
		return false
	}
	if len(ref.tasks) != len(comp.tasks) {
		return false
	}
	for ind, _ := range ref.tasks {
		if ref.tasks[ind].ID != comp.tasks[ind].ID {
			return false
		}
		if ref.tasks[ind].Text != comp.tasks[ind].Text {
			return false
		}
		if ref.tasks[ind].Done != comp.tasks[ind].Done {
			return false
		}
	}
	return true
}

func TestEmptyText(t *testing.T) {
	s := NewService()
	emptyText := ""
	_, err := s.Add(emptyText)
	if !errors.Is(err, ErrEmptyText) {
		t.Fatalf("expected %v error, got %v", ErrEmptyText, err)
	}
}

func TestNotFoundDone(t *testing.T) {
	s := NewService()
	err := s.Done(1)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected %v error, got %v", ErrNotFound, err)
	}
}

func TestNotFoundDelete(t *testing.T) {
	s := NewService()
	err := s.Delete(1)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected %v error, got %v", ErrNotFound, err)
	}
}

func TestAdd(t *testing.T) {
	s := NewService()
	reference := Task{
		ID:   0,
		Text: "reference",
		Done: false,
	}
	refService := Service{
		nextID: 1,
		tasks:  map[int]Task{0: reference},
	}
	s.Add("reference")
	if !compare(refService, *s) {
		t.Fatalf("got wrong Mathes")
	}
}

func TestDelete(t *testing.T) {
	s := NewService()
	task0 := Task{
		ID:   0,
		Text: "one",
		Done: false,
	}
	task2 := Task{
		ID:   2,
		Text: "three",
		Done: false,
	}
	refService := Service{
		nextID: 3,
		tasks:  map[int]Task{0: task0, 2: task2},
	}
	s.Add("one")
	s.Add("two")
	s.Add("three")
	s.Delete(1)
	if !compare(refService, *s) {
		t.Fatalf("got wrong Mathes")
	}
}
