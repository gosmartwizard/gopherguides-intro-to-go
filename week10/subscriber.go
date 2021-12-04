package week10

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type subscriber struct {
	Name   string   `json:"name"`
	Topics []string `json:"topics"`
}

type Subscriber struct {
	Name   string
	Topics []string
	Cancel context.CancelFunc
}

func NewSubscriber(name string) *Subscriber {
	s := &Subscriber{}
	s.Name = name

	return s
}

func (s *Subscriber) SubscriberStart(ctx context.Context, topics ...string) (context.Context, error) {

	ctx, cancel := context.WithCancel(ctx)

	s.Cancel = cancel

	s.Topics = make([]string, len(topics))

	copy(s.Topics, topics)

	return ctx, nil
}

func (s *Subscriber) Listen(ctx context.Context, ch chan Article) {
	go s.readArticle(ctx, ch)
}

func (s *Subscriber) readArticle(ctx context.Context, ch chan Article) {

	for {
		select {
		case <-ctx.Done():
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

func GetSubscribers(fileName string) ([]subscriber, error) {

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var subscribers []subscriber

	err = json.Unmarshal(fileBytes, &subscribers)

	if err != nil {
		return nil, err
	}

	return subscribers, nil
}
