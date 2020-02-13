package query

import (
	"strings"
)

// From from
type From struct {
	tables []string
}

// Set set
func (f *From) Set(str *string) Token {
	f.tables = strings.Split(*str, ",")
	return f
}

// Get Select
func (f *From) Get() Token {
	return f
}
