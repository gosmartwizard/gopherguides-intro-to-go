package week09

import (
	"context"
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

	s.cancel = cancel

	s.categories = make([]string, len(categories))

	copy(s.categories, categories)

	return ctx, nil
}

func (s *MockSource) PublishArticles(ctx context.Context) {

	go func(ctx context.Context) {
		<-ctx.Done()

		s.SourceStop()
	}(ctx)

	s.RLock()
	defer s.RUnlock()

	/* for {

		article := Article{}
		article.source = "Mock_News_Source"
		article.category = "Sports"
		article.description = "Sachin Tendulkar"

		s.publish(article)

		time.Sleep(time.Millisecond * 5000)
	} */

	for {
		if s.name == "MockSource" {
			article := Article{}
			article.Source = "Mock_News_Source_1"
			article.Category = "Sports"
			article.Description = "Sachin Tendulkar"

			s.publish(article)
		}

		time.Sleep(time.Millisecond * 5000)
	}
}

func (s *MockSource) publish(article Article) {
	s.News <- article
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

		s.cancel()

		s.closed = true

		if s.News != nil {
			close(s.News)
		}
	})
}
