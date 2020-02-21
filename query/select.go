package query

import (
	"strings"
)

// Select sel
type Select struct {
	columns []string
}

// Set set
func (s *Select) Set(str *string) Token {
	s.columns = strings.Split(*str, ",")
	return s
}

// Get Select
func (s *Select) Get() Token {
	return s
}

func (s *Select) foundColumns() {}
