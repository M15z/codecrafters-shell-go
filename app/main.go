package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readPrompt() string {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	return command

}

func main() {
	for {
		command := readPrompt()
		command = strings.TrimSpace(command)
		if command == "exit" {
			break
		}

		fmt.Println(command[:len(command)] + ": command not found")
	}
}
