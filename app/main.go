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

func handleEcho(words []string) {
	for _, word := range words[1:] {
		fmt.Print(word + " ")
	}
	fmt.Println()
}

func main() {
	for {
		command := readPrompt()
		if command == "exit" {
			break
		}

		command = strings.TrimSpace(command)

		words := strings.Fields(command)
		if words[0] == "echo" {
			handleEcho(words)
		} else {
			fmt.Printf("%s: command not found\n", command)
		}

	}
}
