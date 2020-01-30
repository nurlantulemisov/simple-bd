package query

import (
	"errors"
	"fmt"
	"strings"
	"regexp"
)

// Query query
type Query struct {
	Sel Select
	From From
	Insert Insert
}
 // Setters set
func (q *Query) Setters(str string) {
	//var columns, from string
	var columns, err = q.Parsing(`(?m)(?:select)(.*?)(?:from)`, str)

	if err == nil {
		q.Sel.SetSelect(columns)
	}

	var columnsInsert, errInsert = q.Parsing(`(?m)(?:insert)(.*?)(?:value)`, str)

	if errInsert == nil {
		q.Insert.Columns = strings.Split(columnsInsert, ",")

		var values, errValue = q.Parsing(`(?m)(?:value)(.*?)(?:from)`, str)

		if errValue == nil {
			q.Insert.Values = strings.Split(values, ",")
		}
	}

	var tables, errTable = q.Parsing(`(?m)from(.*)(?:if|;)`, str)

	if errTable == nil {
		q.From.tables = strings.Split(tables, ",")
	}
}

// Getters get
func (q *Query) Getters() {
	fmt.Println(q.Sel.columns, " + ", q.From.tables)
	fmt.Println(q.Insert.Columns, " + ", q.Insert.Values, " + ", q.From.tables)
}

// Parsing parsing
func (q *Query) Parsing(rex string, str string) (string, error) {
	var re = regexp.MustCompile(rex)
	var columns = re.FindAllStringSubmatch(str, -1)

	if len(columns) > 0 {
		return columns[0][1], nil
	}
	return "", errors.New("not found element")
}