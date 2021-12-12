package main

import "testing"

func Test_Subscriber_GoldenPath(t *testing.T) {
	t.Parallel()

	s := NewSubscriber("Subscriber1")

	s.SubscriberStart("Sports")

	ch := make(chan Article)

	s.Listen(ch)

	article := Article{}
	article.Source = "News_Source_1"
	article.Category = "Sports"
	article.Description = "Sachin Tendulkar"
	ch <- article

	s.Cancel()

	if ch != nil {
		close(ch)
	}
}

func Test_Subscriber_ChannelClose(t *testing.T) {
	t.Parallel()

	s := NewSubscriber("Subscriber2")

	s.SubscriberStart("Movies")

	ch := make(chan Article)

	s.Listen(ch)

	article := Article{}
	article.Source = "News_Source_2"
	article.Category = "Movies"
	article.Description = "Avengers"
	ch <- article

	if ch != nil {
		close(ch)
	}

	s.Cancel()
}
