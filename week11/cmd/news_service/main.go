package main

import (
	"fmt"
	"os"
	"week11/cmd/news_service/cli"
)

func main() {

	app := &cli.App{}

	err := app.Main(os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
