package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"query"
	reader "reader-console"
)

func main() {
	messages := make(chan string)
	go reader.Reader(messages)

	str := <-messages

	var q query.Query
	q.Setters(str)
	q.Getters()

	if len(q.Insert.Values) > 0 {
		q.Insert.From = &q.From
		err := q.Insert.WriteInsert()

		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(q)
}

// WatchFile watch file in dir
func WatchFile(q *query.Query) {
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
