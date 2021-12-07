package cli

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gosmartwizard/gopherguides-intro-to-go/week11/source"
)

// App is the CLI application for the `notes` package
type App struct {
}

// snippet: main

// Main is the entry point for the CLI application
func (app *App) Main(ctx context.Context, pwd string, args []string) error {
	if app == nil {
		return fmt.Errorf("app is nil")
	}

	//TODO
	/* if len(args) == 0 {
		return app.Usage(app.Stdout())
	} */

	if len(os.Args) < 2 {
		// TODO
		fmt.Println("expected 'stream' or 'read' or 'clear' subcommands")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
	case "stream":
		HandleStream()
	case "read":
		return HandleRead()
	case "clear":
		HandleClear()
	default:
		// TODO
	}
	return nil
}

func (app *App) Usage(w io.Writer) error {
	fmt.Fprintln(w, "Usage: notebook <command> [options] [<args>...]")
	fmt.Fprintln(w, "---------------")

	// TODO: print sub-commands

	return nil
}

func HandleStream() {

}

func HandleRead() error {

	var BackupFile string
	var JSON bool
	var OutputFile string

	readCmd := flag.NewFlagSet("read", flag.ExitOnError)

	readCmd.BoolVar(&JSON, "j", false, "output in json format")
	readCmd.StringVar(&BackupFile, "f", "/tmp/news.json", "location of the backupfile")
	readCmd.StringVar(&OutputFile, "o", "", "output results to a file")

	readCmd.Parse(os.Args[2:])

	args := os.Args[1:]
	fmt.Printf("args before parsing: %#v\n", args)

	args = readCmd.Args()
	fmt.Printf("args after parsing: %#v\n", args)

	fmt.Printf("BackupFile : %v\n", BackupFile)
	fmt.Printf("JSON : %t\n", JSON)
	fmt.Printf("OutPutFile : %v\n", OutputFile)

	if len(args) == 0 {
		return fmt.Errorf("ID numbers are not provided")
	}

	for _, id := range args {

		id, _ := strconv.Atoi(id)

		if id <= 0 {
			return fmt.Errorf("Id : %#v is not valid", id)
		}
	}

	articles, err := getArticles(BackupFile, args)

	if err != nil {
		return err
	}

	if len(OutputFile) > 0 {
		saveArticlesInOutputFile(OutputFile, articles)
	}

	if JSON {
		json.NewEncoder(os.Stdout).Encode(articles)
	}

	return nil
}

func getArticles(backupFile string, ids []string) ([]source.Article, error) {

	var articles []source.Article

	fileBytes, err := ioutil.ReadFile(backupFile)

	if err != nil {
		return articles, err
	}

	var idarticles map[string]source.Article

	err = json.Unmarshal(fileBytes, &idarticles)

	if err != nil {
		return articles, err
	}

	for _, id := range ids {

		article, ok := idarticles[id]

		if !ok {
			return articles, fmt.Errorf("Id : %#v doesn't exist", id)
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func saveArticlesInOutputFile(outputFileLocation string, articles []source.Article) error {

	fileBytes, err := json.Marshal(articles)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFileLocation, fileBytes, 0644)

	if err != nil {
		return nil
	}

	return nil
}

func HandleClear() {

}
