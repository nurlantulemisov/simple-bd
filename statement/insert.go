package statement

// Insert represents a SQL INSERT statement.
type Insert struct {
	Fields    map[string]string
	TableName string
}

// SetTable sets name table
func (i *Insert) SetTable(table string) {
	i.TableName = table
}
