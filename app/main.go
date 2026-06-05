package main

import (
	"bufio"
	"fmt"
	"os"
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
	command := readPrompt()
	fmt.Println(command[:len(command)-1] + ": command not found")
}
