package dedup

import (
	"sort"
	"sync"
)

// Structure for tracking strings we have already seen.
type Seen struct {
	seen map[string]bool
	mu   sync.Mutex
}

// Mark a key as visited and return, whether it has been seen already.
func (s *Seen) Visit(key string) (seen bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.seen == nil {
		s.seen = make(map[string]bool)
	}

	if !s.seen[key] {
		s.seen[key] = true
		return false
	}
	return true
}

// Is key already known?
func (s *Seen) Known(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.seen == nil {
		s.seen = make(map[string]bool)
	}

	if !s.seen[key] {
		return false
	}

	return s.seen[key]
}

// Return all visited keys
func (s *Seen) visitedKeys() (keys []string) {
	if s.seen == nil {
		return nil
	}

	for key := range s.seen {
		keys = append(keys, key)
	}
	return keys
}

func (s *Seen) VisitedKeys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.visitedKeys()
}

func (s *Seen) VisitedKeysSorted() []string {
	s.mu.Lock()
	keys := s.visitedKeys()
	s.mu.Unlock()
	sort.Strings(keys)
	return keys
}
