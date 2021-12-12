package week11

import (
	"context"
	"testing"
	"time"
)

func Test_MockSource_Sports_Category(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	mockSource := NewMockSource("MockSource", 10)

	mockSource.SourceStart("Sports")

	go mockSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	cancel()
}

func Test_MockSource_Tech_Category(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	mockSource := NewMockSource("MockSource", 10)

	mockSource.SourceStart("Tech")

	go mockSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	cancel()
}

func Test_MockSource_Movies_Category(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	mockSource := NewMockSource("MockSource", 10)

	mockSource.SourceStart("Movies")

	go mockSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	cancel()
}

func Test_MockSource_Politics_Category(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	mockSource := NewMockSource("MockSource", 10)

	mockSource.SourceStart("Politics")

	go mockSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	cancel()
}

func Test_MockSource_Music_Category(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	mockSource := NewMockSource("MockSource", 10)

	mockSource.SourceStart("Music")

	go mockSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	cancel()
}

func Test_MockSource_Start_Stop(t *testing.T) {
	t.Parallel()

	mockSource := NewMockSource("MockSource", 10)

	mockSource.SourceStart("Sports", "Tech", "Movies")

	mockSource.SourceStop()

	mockSource.SourceStop()
}

func Test_MockSource_MultipleCategory_WithTimeout(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()

	ctx, cancel := context.WithTimeout(rootCtx, 15*time.Second)

	defer cancel()

	mockSource := NewMockSource("MockSource_5", 10)

	mockSource.SourceStart("Sports", "Tech", "Movies")

	go mockSource.PublishArticles(ctx)

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
