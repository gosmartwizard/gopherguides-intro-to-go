package week10

import (
	"context"
	"testing"
	"time"
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

	mockSource := NewMockSource("MockSource_2")

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

	mockSource := NewMockSource("MockSource_3")

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

	mockSource := NewMockSource("MockSource_4")

	_, err := mockSource.SourceStart(ctx, "Sports", "Tech", "Movies")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

	mockSource.SourceStop()

	mockSource.SourceStop()
}

func Test_MockSource_MultipleCategory(t *testing.T) {
	t.Parallel()

	rootCtx := context.Background()

	ctx, cancel := context.WithTimeout(rootCtx, 30*time.Second)

	defer cancel()

	mockSource := NewMockSource("MockSource_5")

	_, err := mockSource.SourceStart(ctx, "Sports", "Tech", "Movies")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

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
