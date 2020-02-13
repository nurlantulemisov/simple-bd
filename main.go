package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"query"
	"queue"
	reader "reader-console"
	"regexp"
)

func main() {
	messages := make(chan string)
	go reader.Reader(messages)

	str := <-messages

	setters(str)
	// q.Setters(str)
	// q.Getters()

	// if len(q.Insert.Values) > 0 {
	// 	q.Insert.From = &q.From
	// 	err := q.Insert.WriteInsert()

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
}

// WatchFile watch file in dir
func WatchFile() {
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

// Setters set
func setters(str string) {
	var columns, _ = parsing(`(?m)(?:select)(.*?)(?:from)`, &str)
	var queueType = new(queue.Queue)
	var selectCommand = new(query.Select)

	createCommand(&str, queueType, selectCommand)

	var tables, _ = parsing(`(?m)from(.*)(?:if|;)`, &str)

	if len(tables) > 0 {
		var fromCommand = new(query.From)
		fromCommand.Set(&tables)
		queueType.Push(fromCommand)
	}
	for _, item := range queueType.GetAll() {
		fmt.Println(*item)
	}
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

func createCommand(str *string, queueType *queue.Queue, command *query.Token) error {
	var columns, err = parsing(`(?m)(?:select)(.*?)(?:from)`, str)

	if err == nil && len(columns) > 0 {
		command.Set(&columns)
		queueType.Push(command)

		return nil
	}
	return err
}

// // SetInsert set insert
// func setInsert(str *string) {
// 	var columnsInsert, errInsert = parsing(`(?m)(?:insert)(.*?)(?:value)`, *str)

// 	if errInsert == nil {
// 		q.Insert.Columns = strings.Split(columnsInsert, ",")

// 		var values, errValue = parsing(`(?m)(?:value)(.*?)(?:from)`, *str)

// 		if errValue == nil {
// 			q.Insert.Values = strings.Split(values, ",")
// 		}
// 	}
// }
