package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	//"github.com/gosmartwizard/gopherguides-intro-to-go/week11/pkg"
)

func HandleRead(args []string) error {

	var BackupFile string
	var JSON bool
	var OutputFile string

	readCmd := flag.NewFlagSet("read", flag.ExitOnError)

	readCmd.BoolVar(&JSON, "j", false, "output in json format")
	readCmd.StringVar(&BackupFile, "f", "/tmp/NewsServiceBackup.json", "location of the backupfile")
	readCmd.StringVar(&OutputFile, "o", "", "output results to a file")

	readCmd.Parse(args)

	args = readCmd.Args()

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
		fmt.Printf("Article : %v is from source : %v under category : %v with description : %v \n", id, article.Source, article.Category, article.Description)
	}

	return nil
}

func getArticles(backupFile string, ids []string) (map[int]Article, error) {

	var articles map[int]Article

	ns := NewNewService()

	err := ns.Start()

	if err != nil {
		return articles, err
	}

	articles, err = ns.GetArticlesByIds(backupFile, ids)

	return articles, err
}

func saveArticlesInOutputFile(outputFileLocation string, articles map[int]Article) error {

	if len(articles) == 0 {
		return nil
	}

	bytes, err := json.Marshal(articles)

	if err != nil {
		return err
	}

	file, err := os.Create(outputFileLocation)

	if err != nil {
		return err
	}

	_, err = file.Write(bytes)

	return err
}
