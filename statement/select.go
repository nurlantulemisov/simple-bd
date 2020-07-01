package statement

// Select represents a SQL SELECT statement.
type Select struct {
	Fields    []string
	TableName string
}

// SetTable sets name table
func (s *Select) SetTable(table string) {
	s.TableName = table
}
