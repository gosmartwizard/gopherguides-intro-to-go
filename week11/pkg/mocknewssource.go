package pkg

import (
	"context"
	"fmt"
	"time"
)

func NewMockSource(name string, mockSourceInterval int) *MockSource {
	s := &MockSource{}
	s.name = name
	s.mockSourceInterval = mockSourceInterval
	s.closed = false

	return s
}

func (s *MockSource) SourceStart(categories ...string) (ch chan []Article) {

	s.Lock()
	defer s.Unlock()

	s.categories = make([]string, 0, len(categories))

	s.categories = append(s.categories, categories...)

	s.News = make(chan []Article)

	return s.News
}

func (s *MockSource) PublishArticles(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			s.SourceStop()
			return
		case s.News <- s.GetArticles():
		}
	}
}

func (ms *MockSource) GetArticles() []Article {

	time.Sleep(time.Second * time.Duration(ms.mockSourceInterval))

	ms.RLock()

	var articles []Article
	article := Article{}

	for _, category := range ms.categories {

		switch category {
		case "sports", "Sports":
			article.Description = "Sachin Tendulkar"
		case "tech", "Tech":
			article.Description = "GoLang"
		case "movies", "Movies":
			article.Description = "Avengers"
		case "politics", "Politics":
			article.Description = "Narendra Modi"
		case "music", "Music":
			article.Description = "Mark Bates Beatles"
		default:
			fmt.Println("Invalid category")

		}

		article.Source = ms.name
		article.Category = category

		articles = append(articles, article)
	}

	ms.RUnlock()

	return articles
}

func (s *MockSource) SourceStop() {

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
