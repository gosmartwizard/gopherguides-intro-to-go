package week10

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

func NewFileBasedSource(name string) *FileBasedSource {
	s := &FileBasedSource{}
	s.name = name
	s.News = make(chan Article)
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

	go func(ctx context.Context) {
		<-ctx.Done()

		s.SourceStop()
	}(ctx)

	s.RLock()
	defer s.RUnlock()

	for {

		fileBytes, err := ioutil.ReadFile("./newsarticles/*.json")

		if err != nil {
			panic(err)
		}

		var articles []Article

		err = json.Unmarshal(fileBytes, &articles)

		if err != nil {
			panic(err)
		}

		for _, article := range articles {
			s.publish(article)
		}

		os.Remove("./newsarticles/*.json")

		time.Sleep(time.Millisecond * 60000)
	}
}

func (s *FileBasedSource) publish(article Article) {
	s.News <- article
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
