package week10

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func NewMockSource(name string) *MockSource {
	s := &MockSource{}
	s.name = name
	s.News = make(chan Article)
	s.closed = false

	return s
}

func (s *MockSource) SourceStart(ctx context.Context, categories ...string) (context.Context, error) {

	s.Lock()
	defer s.Unlock()

	ctx, cancel := context.WithCancel(ctx)

	s.Cancel = cancel

	s.categories = make([]string, len(categories))

	copy(s.categories, categories)

	return ctx, nil
}

func (s *MockSource) PublishArticles(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Cancellation in source : %#v \n", s.name)
			s.SourceStop()
			return
		case s.News <- s.getArticle():
			fmt.Println("Article Published")
		}
	}
}

func (ms *MockSource) getArticle() Article {

	//time.Sleep(time.Millisecond * 5000)

	ms.RLock()

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(ms.categories) - 1
	n := rand.Intn(max-min+1) + min

	c := ms.categories[n]

	article := Article{}
	if c == "Sports" {
		article.Source = "Mock_News_Source"
		article.Category = "Sports"
		article.Description = "Sachin Tendulkar"
	} else if c == "Tech" {
		article.Source = "Mock_News_Source"
		article.Category = "Tech"
		article.Description = "GoLang"
	} else if c == "Movies" {
		article.Source = "Mock_News_Source"
		article.Category = "Movies"
		article.Description = "Avengers"
	}
	ms.RUnlock()

	return article
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
