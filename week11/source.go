package week11

import (
	"context"
	"sync"
)

type NewsSourcer interface {
	SourceStart(categories ...string)
	PublishArticles(ctx context.Context)
	SourceStop()
}

type MockSource struct {
	name       string
	categories []string
	News       chan []Article
	sync.RWMutex
	Cancel             context.CancelFunc
	closed             bool
	stopOnce           sync.Once
	mockSourceInterval int
}

type FileBasedSource struct {
	name       string
	filePath   string
	categories []string
	News       chan []Article
	sync.RWMutex
	Cancel            context.CancelFunc
	closed            bool
	stopOnce          sync.Once
	fileBasedInterval int
}

type Article struct {
	Source      string
	Category    string
	Description string
}
