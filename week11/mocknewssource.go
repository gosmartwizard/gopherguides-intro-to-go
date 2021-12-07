package week10

import (
	"context"
	"fmt"
	"time"
)

func NewMockSource(name string) *MockSource {
	s := &MockSource{}
	s.name = name
	s.News = make(chan []Article)
	s.closed = false

	return s
}

func (s *MockSource) SourceStart(interval int, categories ...string) {

	s.Lock()
	defer s.Unlock()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	s.ctx = ctx

	s.Cancel = cancel

	s.categories = make([]string, 0, len(categories))

	copy(s.categories, categories)

	s.mockSourceInterval = interval
}

func (s *MockSource) PublishArticles() {

	for {
		select {
		case <-s.ctx.Done():
			fmt.Printf("Cancellation in source : %#v \n", s.name)
			s.SourceStop()
			return
		case s.News <- s.getArticle():
		}
	}
}

func (ms *MockSource) getArticle() []Article {

	time.Sleep(time.Second * time.Duration(ms.mockSourceInterval))

	ms.RLock()

	var articles []Article
	article := Article{}

	for _, category := range ms.categories {
		if category == "Sports" {
			article.Source = ms.name
			article.Category = category
			article.Description = "Sachin Tendulkar"
		} else if category == "Tech" {
			article.Source = ms.name
			article.Category = category
			article.Description = "GoLang"
		} else if category == "Movies" {
			article.Source = ms.name
			article.Category = category
			article.Description = "Avengers"
		}

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

		s.Cancel()

		s.closed = true

		if s.News != nil {
			close(s.News)
		}
	})
}
