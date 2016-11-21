package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
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
		//Do nothing
	} else {
		commands := strings.Fields(command)
		command = commands[0]
		args := commands[1:]
		//fmt.Println(command, args)
		out, err := exec.Command(command, args...).Output()
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Printf("%s", out)
	}
}
