package main

// snippet: imports
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gosmartwizard/gopherguides-intro-to-go/week11/cmd/news_service/cli"
)

// snippet: imports

// snippet: main
func main() {
	ctx := context.Background()

	// create a context that is cancelled when a SIGINT is received.
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	// get the present working directory. (PWD)
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	app := &cli.App{}

	// call the App's Main method with the context,
	// present working directory, and arguments.
	err = app.Main(ctx, pwd, os.Args[1:])

	// if there was an error, print it and exit with a non-zero status code.
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// wait for the context to be cancelled.
	<-ctx.Done()
}

// snippet: main
