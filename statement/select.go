package statement

// Select represents a SQL SELECT statement.
type Select struct {
	Fields    []string
	TableName string
}
