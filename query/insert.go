package query

import (
	"strings"
	"io"
	"os"
)

// Insert insert
type Insert struct {
	Columns []string
	Values [] string
	From *From
}

// WriteInsert Write to file
func (i *Insert) WriteInsert() error {
	var path string = "./data/" + i.From.tables[0]
	
	file, err := os.Create(path)
    if err != nil {
        return err
    }
	defer file.Close()

	var data string = ""
	for _, column := range i.Columns {
		data += strings.Trim(column, " ") + ";"
	}

	data += "\n"
	for _, value := range i.Values {
		data += strings.Trim(value, " ") + ";"
	}
	data += "\n"

    _, err = io.WriteString(file, data)
    if err != nil {
        return err
    }
    return file.Sync()
}
