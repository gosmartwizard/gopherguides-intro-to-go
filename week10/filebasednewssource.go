package week10

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func NewFileBasedSource(name string, filePath string) *FileBasedSource {
	s := &FileBasedSource{}
	s.name = name
	s.filePath = filePath
	s.News = make(chan []Article)
	s.closed = false

	return s
}

func (s *FileBasedSource) SourceStart(ctx context.Context, categories ...string) (context.Context, error) {

	s.Lock()
	defer s.Unlock()

	ctx, cancel := context.WithCancel(ctx)

	s.Cancel = cancel

	s.categories = make([]string, len(categories))

	copy(s.categories, categories)

	return ctx, nil
}

func (s *FileBasedSource) PublishArticles(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Cancellation in source : %#v \n", s.name)
			s.SourceStop()
			return
		case s.News <- s.GetArticles():
		}
	}
}

func (s *FileBasedSource) GetArticles() []Article {

	time.Sleep(time.Millisecond * 60000)

	var articles []Article

	fileBytes, err := ioutil.ReadFile(s.filePath)

	if err != nil {
		return articles
	}

	err = json.Unmarshal(fileBytes, &articles)

	if err != nil {
		return articles
	}

	os.Remove(s.filePath)

	return articles
}

func (s *FileBasedSource) SourceStop() {

	s.RLock()
	if s.closed {
		s.RUnlock()
		return
	}
	s.RUnlock()

	s.stopOnce.Do(func() {
		s.Lock()
		defer s.Unlock()

		s.Cancel()

		s.closed = true

		if s.News != nil {
			close(s.News)
		}
	})
}
