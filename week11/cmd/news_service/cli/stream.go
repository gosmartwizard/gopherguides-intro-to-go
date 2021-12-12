package cli

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
	"week11/newsservice"
	"week11/source"
)

func HandleStream(args []string) error {

	var BackupFile string
	var JSON bool
	var OutputFile string

	readCmd := flag.NewFlagSet("stream", flag.ExitOnError)

	readCmd.BoolVar(&JSON, "j", false, "output in json format")
	readCmd.StringVar(&BackupFile, "f", "/tmp/NewsServiceBackup.json", "location of the backupfile")
	readCmd.StringVar(&OutputFile, "o", "", "output results to a file")

	readCmd.Parse(args)

	args = readCmd.Args()

	if len(args) == 0 {
		return fmt.Errorf("categories not provided")
	}

	err := getStreamByCategory(BackupFile, args, OutputFile, JSON)

	if err != nil {
		return err
	}

	return nil
}

func getStreamByCategory(backupFile string, categories []string, outputFile string, JSON bool) error {

	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	defer cancel()

	ns := newsservice.NewNewService()

	err := ns.Start()

	if err != nil {
		return err
	}

	filePath := "./newsarticles/newsarticles.json"

	ns.FileBasedRegistration("FileBasedSource1", filePath, categories...)

	go func() {
		for {
			time.Sleep(time.Second * time.Duration(10))

			var articles map[int]source.Article

			articles, _ = ns.GetStreamByCategory(backupFile, categories)

			if len(outputFile) > 0 {
				saveArticlesInOutputFile(outputFile, articles)
				continue
			}

			if JSON {
				json.NewEncoder(os.Stdout).Encode(articles)
				continue
			}

			for id, article := range articles {
				fmt.Printf("Article : %v is from source : %v under category : %v with description : %v \n", id, article.Source, article.Category, article.Description)
			}

		}
	}()

	<-ctx.Done()

	ns.Stop()

	return nil
}
