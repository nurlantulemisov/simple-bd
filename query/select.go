package query

import (
	"strings"
)

// Select sel
type Select struct {
	columns []string
}

// SetSelect set
func (s *Select) SetSelect(str string) {
	s.columns = strings.Split(str, ",")
}