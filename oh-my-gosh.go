package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Echo Shell")

	for {
		fmt.Print("echo > ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		parseCommand(text)
	}

}

func parseCommand(command string) {
	if command == "exit" {
		os.Exit(1)
	} else if command == "" {

	} else {
		fmt.Println(command)
	}
}
