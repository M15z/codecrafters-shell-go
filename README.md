# gosh — A Unix Shell in Go

A minimal Unix shell built from scratch in Go. Supports built-in commands, external program execution, and shell-style quoting.

## Features

- Built-in commands: `echo`, `pwd`, `cd`, `type`, `exit`
- External command execution via `$PATH` lookup
- Quote handling: single quotes, double quotes, and backslash escaping
- Home directory shortcut with `cd ~`

## Build & Run

```bash
go build -o gosh ./app
./gosh
```

## Usage

```
$ echo hello world
hello world

$ echo 'spaces   preserved'
spaces   preserved

$ pwd
/home/mohammed

$ cd ~/projects
$ type echo
echo is a shell builtin

$ type ls
ls is /usr/bin/ls

$ ls -la

$ exit
```

## What I Learned

Built as a Go learning project covering: string manipulation, runes and UTF-8 encoding, OS syscalls (`chdir`, `getwd`), process spawning with `os/exec`, file descriptors, and shell parsing (quoting, escaping).
