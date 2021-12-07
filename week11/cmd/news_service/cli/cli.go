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

	if len(args) == 0 {
		return app.Usage(os.Stdout)
	}

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
	fmt.Fprintln(w, "Usage: new_service <command> [options] [<args>...]")

	fmt.Fprintln(w, "---------------")

	fmt.Fprintln(w, "news_service stream -f /tmp/news.json -j -o ./stream.json sports")

	fmt.Fprintln(w, "news_service read -o ./articles.json -j -f /tmp/news.json 1 2 3")

	fmt.Fprintln(w, "news_service clear -f /tmp/news_service.json")

	return nil
}

func HandleStream() {

}

type Article struct {
	Source      string `json:"Source"`
	Category    string `json:"Category"`
	Description string `json:"Description"`
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

	//args := os.Args[1:]
	//fmt.Printf("args before parsing: %#v\n", args)

	args := readCmd.Args()
	//fmt.Printf("args after parsing: %#v\n", args)

	//fmt.Printf("BackupFile : %v\n", BackupFile)
	//fmt.Printf("JSON : %t\n", JSON)
	//fmt.Printf("OutPutFile : %v\n", OutputFile)

	if len(args) == 0 {
		return fmt.Errorf("ID numbers are not provided")
	}

	for _, id := range args {

		id, _ := strconv.Atoi(id)

		if id <= 0 {
			return fmt.Errorf("id : %#v is not valid", id)
		}
	}

	articles, err := getArticles(BackupFile, args)

	if err != nil {
		return err
	}

	if len(OutputFile) > 0 {
		saveArticlesInOutputFile(OutputFile, articles)
		return nil
	}

	if JSON {
		json.NewEncoder(os.Stdout).Encode(articles)
		return nil
	}

	for _, article := range articles {
		fmt.Printf("Article is from source : %v under category : %v with description : %v \n", article.Source, article.Category, article.Description)
	}

	return nil
}

func getArticles(backupFile string, ids []string) ([]Article, error) {

	var articles []Article

	fileBytes, err := ioutil.ReadFile(backupFile)

	if err != nil {
		return articles, err
	}

	var idarticles map[int]Article

	err = json.Unmarshal(fileBytes, &idarticles)

	if err != nil {
		return articles, err
	}

	//fmt.Printf("idarticles : %#v \n", idarticles)

	for _, id := range ids {

		id, _ := strconv.Atoi(id)

		article, ok := idarticles[id]

		if !ok {
			return articles, fmt.Errorf("article id : %#v doesn't exist", id)
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func saveArticlesInOutputFile(outputFileLocation string, articles []Article) error {

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
