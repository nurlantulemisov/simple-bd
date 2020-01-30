package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Reader read console
func Reader(ch chan string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Shell Console")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("quit", text) == 0 {
			ch <- "Goodbye"
			break
		}

		if text != "" {
			ch <- strings.Replace(text, "-> ", "", -1)
			break
		}
	}
}