package week10

import (
	"testing"
)

func Test_Subscriber_GoldenPath(t *testing.T) {
	t.Parallel()

	s := NewSubscriber("Subscriber1")

	s.SubscriberStart("Sports")

	ch := make(chan Article)

	go s.Listen(ch)

	article := Article{}
	article.Source = "News_Source_1"
	article.Category = "Sports"
	article.Description = "Sachin Tendulkar"
	ch <- article

	s.Cancel()
}

func Test_Subscriber_ChannelClose(t *testing.T) {
	t.Parallel()

	//ctx := context.Background()

	s := NewSubscriber("Subscriber2")

	s.SubscriberStart("Movies")

	ch := make(chan Article)

	go s.Listen(ch)

	article := Article{}
	article.Source = "News_Source_2"
	article.Category = "Movies"
	article.Description = "Avengers"
	ch <- article

	close(ch)

	s.Cancel()
}

func Test_Subscriber_WrongCategory(t *testing.T) {
	t.Parallel()

	s := NewSubscriber("Subscriber3")

	s.SubscriberStart("Sports")

	ch := make(chan Article)

	go s.Listen(ch)

	article := Article{}
	article.Source = "News_Source_3"
	article.Category = "Tech"
	article.Description = "Go Lang"
	ch <- article

	s.Cancel()
}

/* func Test_GetSubscribers_GoldenPath(t *testing.T) {
	t.Parallel()

	subscribers, err := GetSubscribers("./testdata/subscribers.json")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

	exp := 3

	got := len(subscribers)

	if exp != got {
		t.Fatalf("Expected : %#v, got : %#v", exp, got)
	}
}

func Test_GetSubscribers_EmptyFile(t *testing.T) {
	t.Parallel()

	_, err := GetSubscribers("./testdata/subscribers_empty.json")

	if err == nil {
		t.Fatalf("Expected : %#v, got : nil", err)
	}
}

 func Test_GetSubscribers_WrongData(t *testing.T) {
	t.Parallel()

	_, err := GetSubscribers("./testdata/subscribers_wrongdata.json")

	if err == nil {
		t.Fatalf("Expected : %#v, got : nil", err)
	}
}

func Test_GetSubscribers_FileNotExists(t *testing.T) {
	t.Parallel()

	_, err := GetSubscribers("./testdata/subscribers_filenotexist.json")

	if err == nil {
		t.Fatalf("Expected : %#v, got : nil", err)
	}
} */
