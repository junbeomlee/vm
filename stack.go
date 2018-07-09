package vm

import (
	"errors"
	"sync"
)

type Stack struct {
	lock  sync.Mutex
	items [][]uint8
}

func (s *Stack) Push(v []uint8) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append(s.items, v)
}

func (s *Stack) Pop() ([]uint8, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.items)

	if l == 0 {
		return []uint8{}, errors.New("Empty Stack")
	}

	res := s.items[l-1]
	s.items = s.items[:l-1]
	return res, nil
}
