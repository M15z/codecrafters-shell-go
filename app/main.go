package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

func isBuiltIn(arg string) bool {
	builtin := []string{"echo", "exit", "type"}

	for _, b := range builtin {
		if b == arg {
			return true
		}
	}

	return false
}

func handleType(words []string) {
	args := words[1]

	if isBuiltIn(args) {
		fmt.Printf("%s is a shell builtin\n", args)
		return
	} else if path, _ := exec.LookPath(args); path != "" {
		fmt.Println(args + " is " + path)
		return
	}

	fmt.Printf("%s: not found\n", args)
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
