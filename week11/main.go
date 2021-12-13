package main

import (
	"fmt"
	"os"
)

func main() {

	app := &App{}

	err := app.Main(os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
