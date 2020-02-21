package query

import (
	"strings"
)

// Insert insert
type Insert struct {
	columns []string
}

// Set insert
func (i *Insert) Set(str *string) Token {
	i.columns = strings.Split(*str, ",")
	return i
}

// Get insert model
func (i *Insert) Get() Token {
	return i
}

//GetColumns columns
func (i *Insert) GetColumns() *[]string {
	return &i.columns
}

// Value for values
type Value struct {
	values []string
}

//GetValues columns
func (v *Value) GetValues() *[]string {
	return &v.values
}

// Set value
func (v *Value) Set(str *string) Token {
	v.values = strings.Split(*str, ",")
	return v
}

// Get value model
func (v *Value) Get() Token {
	return v
}
