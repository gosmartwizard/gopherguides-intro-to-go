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
	//"github.com/gosmartwizard/gopherguides-intro-to-go/week10/source"
)

type App struct {
}

func (app *App) Main(ctx context.Context, pwd string, args []string) error {
	if app == nil {
		return fmt.Errorf("app is nil")
	}

	if len(args) == 0 {
		return getNewserviceStats()
	}

	if os.Args[1] == "-h" {
		return app.Usage(os.Stdout)
	}

	switch os.Args[1] {
	case "stream":
		return HandleStream()
	case "read":
		return HandleRead()
	case "clear":
		HandleClear()
	default:
		return fmt.Errorf("%v command not supported", os.Args[1])
	}
	return nil
}

func getNewserviceStats() error {
	//TODO
	fmt.Println("news service stats need to do ")
	return nil
}

func (app *App) Usage(w io.Writer) error {
	fmt.Fprintln(w, "Usage: news_service <command> [options] [<args>...]")

	fmt.Fprintln(w, "---------------")

	fmt.Fprintln(w, "news_service stream -f /tmp/news.json -j -o ./stream.json sports")

	fmt.Fprintln(w, "news_service read -o ./articles.json -j -f /tmp/news.json 1 2 3")

	fmt.Fprintln(w, "news_service clear -f /tmp/news_service.json")

	return nil
}

func HandleStream() error {

	var BackupFile string
	var JSON bool
	var OutputFile string

	readCmd := flag.NewFlagSet("stream", flag.ExitOnError)

	readCmd.BoolVar(&JSON, "j", false, "output in json format")
	readCmd.StringVar(&BackupFile, "f", "/tmp/news.json", "location of the backupfile")
	readCmd.StringVar(&OutputFile, "o", "", "output results to a file")

	readCmd.Parse(os.Args[2:])

	args := readCmd.Args()

	if len(args) == 0 {
		return fmt.Errorf("categories are not provided")
	}

	articles, err := getStreamByCategory(BackupFile, args)

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

	for id, article := range articles {
		fmt.Printf("Article : %v is from source : %v under category : %v with description : %v \n", args[id], article.Source, article.Category, article.Description)
	}

	return nil
}

func getStreamByCategory(backupFile string, categories []string) ([]Article, error) {

	var articles []Article

	//TODO
	//The stream should be cancellable through a ctrl-c event.
	return articles, nil
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

	args := readCmd.Args()

	if len(args) == 0 {
		return fmt.Errorf("ID numbers are not provided")
	}

	for _, id := range args {

		id, err := strconv.Atoi(id)

		if err != nil {
			return err
		}

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

	for id, article := range articles {
		fmt.Printf("Article : %v is from source : %v under category : %v with description : %v \n", args[id], article.Source, article.Category, article.Description)
	}

	return nil
}

func getArticles(backupFile string, ids []string) ([]Article, error) {

	//TODO
	articles := []Article{}

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

func HandleClear() error {

	var BackupFile string

	readCmd := flag.NewFlagSet("clear", flag.ExitOnError)
	readCmd.StringVar(&BackupFile, "f", "/tmp/news.json", "location of the backupfile")

	readCmd.Parse(os.Args[2:])

	clear(BackupFile)

	return nil
}

func clear(backupFile string) {

	//TODO
	os.Remove(backupFile)
}
