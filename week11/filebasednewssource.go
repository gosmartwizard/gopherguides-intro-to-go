package week11

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

func NewFileBasedSource(name string, filePath string, interval int) *FileBasedSource {
	s := &FileBasedSource{}
	s.name = name
	s.filePath = filePath
	s.fileBasedInterval = interval
	s.closed = false

	return s
}

func (s *FileBasedSource) SourceStart(categories ...string) (ch chan []Article) {

	s.Lock()
	defer s.Unlock()

	s.categories = make([]string, 0, len(categories))

	s.categories = append(s.categories, categories...)

	s.News = make(chan []Article)

	return s.News

}

func (s *FileBasedSource) PublishArticles(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			s.SourceStop()
			return
		case s.News <- s.GetArticles():
		}
	}
}

func (s *FileBasedSource) GetArticles() []Article {

	time.Sleep(time.Second * time.Duration(s.fileBasedInterval))

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

		s.closed = true

		if s.News != nil {
			close(s.News)
		}
	})
}
