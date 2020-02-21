package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"query"
	"queue"
	reader "reader-console"
	"regexp"
	"strings"
)

func main() {
	messages := make(chan string)
	go reader.Reader(messages)

	str := <-messages

	var queueType = new(queue.Queue)
	setters(str, queueType)

	for {
		var command = queueType.Pop()

		if command != nil {
			typeValue := *command

			switch typeValue.(type) {
			case *query.Select:
				fmt.Println("select")
			case *query.From:
				fmt.Println("from")
			default:
				fmt.Println("type unknown")
			}
		} else {
			fmt.Println("Element isn't found")
			break
		}
	}
}

// Setters set
func setters(str string, queueType *queue.Queue) {
	var selectCommand = new(query.Select)

	var err = createCommand(&str, `(?m)(?:select)(.*?)(?:from)`, queueType, selectCommand)

	if err != nil {
		var insertCommand = new(query.Insert)
		var valueCommand = new(query.Value)

		var errInsert = createCommand(&str, `(?m)(?:insert)(.*?)(?:value)`, queueType, insertCommand)
		var errValueInsert = createCommand(&str, `(?m)(?:value)(.*?)(?:from)`, queueType, valueCommand)

		if errInsert != nil || errValueInsert != nil {
			panic(errors.New("You forget command"))
		}
	}

	var fromCommand = new(query.From)

	err = createCommand(&str, `(?m)from(.*)(?:if|;)`, queueType, fromCommand)

	if err != nil {
		panic(errors.New("You forget from command"))
	}
}

func createCommand(str *string, rex string, queueType *queue.Queue, command query.Token) error {
	var columns, err = parsing(rex, str)

	if err == nil && len(columns) > 0 {
		command.Set(&columns)
		queueType.Push(&command)

		return nil
	}
	return err
}

// Parsing parsing
func parsing(rex string, str *string) (string, error) {
	var re = regexp.MustCompile(rex)
	var columns = re.FindAllStringSubmatch(*str, -1)

	if len(columns) > 0 {
		return columns[0][1], nil
	}
	return "", errors.New("not found element")
}

//insert Write to file
func write(i *query.Insert, v *query.Value) error {
	var path string = "./data/"

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var data string = ""
	for _, column := range *i.GetColumns() {
		data += strings.Trim(column, " ") + ";"
	}

	data += "\n"
	for _, value := range *v.GetValues() {
		data += strings.Trim(value, " ") + ";"
	}
	data += "\n"

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
