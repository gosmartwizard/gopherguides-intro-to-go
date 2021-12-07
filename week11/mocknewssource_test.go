package week10

import (
	"context"
	"testing"
	"time"
)

func Test_MockSource_Sports_Category(t *testing.T) {
	t.Parallel()

	mockSource := NewMockSource("MockSource_1")

	mockSource.SourceStart(10, "Sports")

	go mockSource.PublishArticles()

	select {
	case <-mockSource.ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	mockSource.Cancel()
}

func Test_MockSource_Tech_Category(t *testing.T) {
	t.Parallel()

	mockSource := NewMockSource("MockSource_2")

	mockSource.SourceStart(10, "Tech")

	go mockSource.PublishArticles()

	select {
	case <-mockSource.ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	mockSource.Cancel()
}

func Test_MockSource_Movies_Category(t *testing.T) {
	t.Parallel()

	mockSource := NewMockSource("MockSource_3")

	mockSource.SourceStart(10, "Movies")

	go mockSource.PublishArticles()

	select {
	case <-mockSource.ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	mockSource.Cancel()
}

func Test_MockSource_Start_Stop(t *testing.T) {
	t.Parallel()

	mockSource := NewMockSource("MockSource_4")

	mockSource.SourceStart(10, "Sports", "Tech", "Movies")

	mockSource.SourceStop()

	mockSource.SourceStop()
}

func Test_MockSource_MultipleCategory_WithTimeout(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()

	ctx, cancel := context.WithTimeout(rootCtx, 30*time.Second)

	defer cancel()

	mockSource := NewMockSource("MockSource_5")

	mockSource.SourceStart(10, "Sports", "Tech", "Movies")

	go mockSource.PublishArticles()

	select {
	case <-rootCtx.Done():
		mockSource.SourceStop()
	case <-ctx.Done():
	}

	exp := context.DeadlineExceeded.Error()

	if exp != ctx.Err().Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, ctx.Err().Error())
	}
}
