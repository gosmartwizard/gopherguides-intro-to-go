package main

import (
	"fmt"
	"os"
	//"github.com/gosmartwizard/gopherguides-intro-to-go/week11/cmd/newservice/cli"
)

func main() {

	app := &App{}

	err := app.Main(os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
