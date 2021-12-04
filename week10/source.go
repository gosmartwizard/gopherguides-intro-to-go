package week10

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"sync"
)

type NewsSourcer interface {
	SourceStart(ctx context.Context, categories ...string)
	PublishArticles(ctx context.Context)
	SourceStop()
}

type source struct {
	Name       string   `json:"name"`
	FilePath   string   `json:"filepath"`
	Categories []string `json:"categories"`
}

type MockSource struct {
	name       string
	categories []string
	News       chan []Article
	sync.RWMutex
	Cancel   context.CancelFunc
	closed   bool
	stopOnce sync.Once
}

type FileBasedSource struct {
	name       string
	filePath   string
	categories []string
	News       chan []Article
	sync.RWMutex
	Cancel   context.CancelFunc
	closed   bool
	stopOnce sync.Once
}

type Article struct {
	Source      string
	Category    string
	Description string
}

func GetSources() ([]source, error) {

	fileBytes, err := ioutil.ReadFile("./sources.json")

	if err != nil {
		panic(err)
	}

	var sources []source

	err = json.Unmarshal(fileBytes, &sources)

	if err != nil {
		panic(err)
	}

	return sources, nil
}
