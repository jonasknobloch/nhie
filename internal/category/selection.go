package category

import "math/rand"

type Selection struct {
	m map[Category]struct{}
	s []Category
}

func NewSelection() *Selection {
	return &Selection{
		m: make(map[Category]struct{}),
		s: make([]Category, 0),
	}
}

func (s *Selection) Add(category Category) {
	if s.Has(category) {
		return
	}

	s.m[category] = struct{}{}
	s.s = append(s.s, category)
}

func (s *Selection) Has(category Category) bool {
	_, ok := s.m[category]

	return ok
}

func (s *Selection) Random() Category {
	if len(s.s) == 0 {
		return Harmless
	}

	return s.s[rand.Intn(len(s.s))]
}

func (s *Selection) Empty() bool {
	return len(s.m) == 0
}
