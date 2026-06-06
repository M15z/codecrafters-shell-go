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

func handleType(words []string) {
	builtin := []string{"echo", "exit", "type"}
	args := words[1]

	for _, b := range builtin {
		if args == b {
			fmt.Printf("%s is a shell builtin\n", args)
			return
		}
	}

	fmt.Print("%s: not found\n", args)
}

func main() {
	for {
		command := readPrompt()
		command = strings.TrimSpace(command)
		if command == "exit" {
			break
		}

		words := strings.Fields(command)
		if words[0] == "echo" {
			handleEcho(words)
		} else if words[0] == "type" {
			handleType(words)
		} else {
			fmt.Printf("%s: command not found\n", command)
		}

	}
}
