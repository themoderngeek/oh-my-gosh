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
		parseCommand(text)
	}

}

func parseCommand(command string) {
	if command == "exit\n" {
		os.Exit(1)
	} else if command == "\n" {

	} else {
		fmt.Print(command)
	}
}
