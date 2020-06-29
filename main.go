package main

import (
	"bufio"
	"fmt"
	"os"
	"scanner"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	parser := scanner.NewParser(strings.NewReader(text))
	stmt, _ := parser.Parse()
	fmt.Println(stmt)
}
