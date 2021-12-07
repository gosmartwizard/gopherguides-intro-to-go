package week10

import (
	"context"
	"fmt"
)

type Subscriber struct {
	Name   string
	Topics []string
	Cancel context.CancelFunc
	ctx    context.Context
}

func NewSubscriber(name string) *Subscriber {
	s := &Subscriber{}
	s.Name = name

	return s
}

func (s *Subscriber) SubscriberStart(topics ...string) {

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	s.ctx = ctx

	s.Cancel = cancel

	s.Topics = make([]string, 0, len(topics))

	copy(s.Topics, topics)
}

func (s *Subscriber) Listen(ch chan Article) {
	go s.readArticle(ch)
}

func (s *Subscriber) readArticle(ch chan Article) {

	for {
		select {
		case <-s.ctx.Done():
			fmt.Printf("Cancellation in Subscriber : %v \n", s.Name)
			return
		case article, ok := <-ch:
			if !ok {
				fmt.Printf("Channel closed in Subscriber : %v \n", s.Name)
				continue
			}

			subscribed := false
			for _, category := range s.Topics {
				if article.Category == category {
					subscribed = true
					break
				} else {
					continue
				}
			}

			if subscribed {
				fmt.Printf("Article : %#v \n", article)
			} else {
				fmt.Printf("Subscriber : %#v not subscribed to this category : %#v \n", s.Name, article.Category)
			}
		}
	}
}
