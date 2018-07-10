package vm

import (
	"errors"
	"sync"
)

type Stack struct {
	lock  sync.RWMutex
	items []Hexable
}

func NewStack() Stack {
	return Stack{
		lock:  sync.RWMutex{},
		items: make([]Hexable, 0),
	}
}

func (s *Stack) Push(v Hexable) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append(s.items, v)
}

func (s *Stack) Pop() (Hexable, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.items)

	if l == 0 {
		return nil, errors.New("Empty Stack")
	}

	res := s.items[l-1]
	s.items = s.items[:l-1]

	return res, nil
}

func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.items)
}
