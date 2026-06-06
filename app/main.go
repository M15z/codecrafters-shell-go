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
	return strings.TrimSpace(command)
}

func isBuiltIn(arg string) bool {
	builtins := []string{"echo", "exit", "type", "pwd"}
	for _, b := range builtins {
		if b == arg {
			return true
		}
	}
	return false
}

func handleEcho(words []string) {
	fmt.Println(strings.Join(words[1:], " "))
}

func handleType(words []string) {
	arg := words[1]
	if isBuiltIn(arg) {
		fmt.Printf("%s is a shell builtin\n", arg)
		return
	}
	if path, _ := exec.LookPath(arg); path != "" {
		fmt.Printf("%s is %s\n", arg, path)
		return
	}
	fmt.Printf("%s: not found\n", arg)
}

func handlePwd() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
}

func handleExternal(words []string) {
	cmd := exec.Command(words[0], words[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func handleCd(path string) {
	if path == "~" {
		home, _ := os.UserHomeDir()
		path = home
	}

	err := os.Chdir(path)

	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path)
	}

}

func dispatch(command string) {
	words := strings.Fields(command)
	if len(words) == 0 {
		return
	}

	switch words[0] {
	case "echo":
		handleEcho(words)
	case "type":
		handleType(words)
	case "pwd":
		handlePwd()
	case "cd":
		handleCd(words[1])
	default:
		if _, err := exec.LookPath(words[0]); err == nil {
			handleExternal(words)
		} else {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}

func main() {
	for {
		command := readPrompt()
		if command == "exit" {
			break
		}
		dispatch(command)
	}
}

