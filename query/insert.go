package query

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Insert insert
type Insert struct {
	Columns []string
	Values  []string
}

// WriteInsert Write to file
func (i *Insert) WriteInsert() error {
	var path string = "./data/"

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

// Set insert
func (i *Insert) Set(str string) {
	i.Columns = strings.Split(str, ",")
	i.Values = strings.Split(str, ",")
}

// Get insert
func (i *Insert) Get() {
	fmt.Println(i.Columns)
	fmt.Println(i.Values)
}
