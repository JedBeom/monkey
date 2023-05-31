package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current() // returns the current user
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Type your commands here\n")
	repl.Start(os.Stdin, os.Stdout)
}
