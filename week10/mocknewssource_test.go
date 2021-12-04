package week10

import (
	"context"
	"testing"
)

func Test_MockSource_Sports_Category(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mockSource := NewMockSource("MockSource_1")

	ctx, err := mockSource.SourceStart(ctx, "Sports")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

	go mockSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	mockSource.Cancel()
}

func Test_MockSource_Tech_Category(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mockSource := NewMockSource("MockSource_1")

	ctx, err := mockSource.SourceStart(ctx, "Tech")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

	go mockSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	mockSource.Cancel()
}

func Test_MockSource_Movies_Category(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mockSource := NewMockSource("MockSource_1")

	ctx, err := mockSource.SourceStart(ctx, "Movies")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

	go mockSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-mockSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	mockSource.Cancel()
}

func Test_MockSource_Start_Stop(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mockSource := NewMockSource("MockSource_1")

	_, err := mockSource.SourceStart(ctx, "Sports", "Tech", "Movies")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

	mockSource.SourceStop()

	mockSource.SourceStop()
}
