package week10

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func generateMultipleCategoryData(filePath string) error {

	var articles []Article
	article := Article{}

	article.Source = "FileBasedSource"
	article.Category = "Sports"
	article.Description = "Brian Lara"

	articles = append(articles, article)

	article.Source = "FileBasedSource"
	article.Category = "Movies"
	article.Description = "Dear Comrade"

	articles = append(articles, article)

	article.Source = "FileBasedSource"
	article.Category = "Tech"
	article.Description = "Kubernetes"

	articles = append(articles, article)

	fileBytes, err := json.Marshal(articles)

	if err != nil {
		return err
	}

	os.Remove(filePath)

	ioutil.WriteFile(filePath, fileBytes, 0644)

	return nil
}

func generateSingleCategoryData(filePath string) error {

	var articles []Article
	article := Article{}

	article.Source = "FileBasedSource"
	article.Category = "Sports"
	article.Description = "Brian Lara"

	articles = append(articles, article)

	fileBytes, err := json.Marshal(articles)

	if err != nil {
		return err
	}

	os.Remove(filePath)

	ioutil.WriteFile(filePath, fileBytes, 0644)

	return nil
}
func Test_FileBasedSource_MultipleCategory(t *testing.T) {
	t.Parallel()

	generateMultipleCategoryData("./testdata/newsarticles.json")

	ctx := context.Background()

	fileBasedSource := NewFileBasedSource("FileBasedSource", "./testdata/newsarticles.json")

	ctx, err := fileBasedSource.SourceStart(ctx, "Sports")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

	go fileBasedSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-fileBasedSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	fileBasedSource.Cancel()
}

func Test_FileBasedSource_Sports_Category(t *testing.T) {
	t.Parallel()

	generateSingleCategoryData("./testdata/newsarticles.json")

	ctx := context.Background()

	fileBasedSource := NewFileBasedSource("FileBasedSource", "./testdata/newsarticles.json")

	ctx, err := fileBasedSource.SourceStart(ctx, "Tech")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

	go fileBasedSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-fileBasedSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	fileBasedSource.Cancel()
}

func Test_FileBasedSource_Start_Stop(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	fileBasedSource := NewFileBasedSource("FileBasedSource", "./testdata/newsarticles.json")

	_, err := fileBasedSource.SourceStart(ctx, "Sports", "Tech", "Movies")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

	fileBasedSource.SourceStop()

	fileBasedSource.SourceStop()
}

func Test_FileBasedSource_WithTimeOut(t *testing.T) {
	t.Parallel()

	generateMultipleCategoryData("./testdata/newsarticles.json")

	rootCtx := context.Background()

	ctx, cancel := context.WithTimeout(rootCtx, 30*time.Second)

	defer cancel()

	fileBasedSource := NewFileBasedSource("FileBasedSource", "./testdata/newsarticles.json")

	_, err := fileBasedSource.SourceStart(ctx, "Sports", "Tech", "Movies")

	if err != nil {
		t.Fatalf("Expected : nil, got : %#v", err)
	}

	go fileBasedSource.PublishArticles(ctx)

	select {
	case <-rootCtx.Done():
		fileBasedSource.SourceStop()
	case <-ctx.Done():
	}

	exp := context.DeadlineExceeded.Error()

	if exp != ctx.Err().Error() {
		t.Fatalf("expected : %#v, got : %#v", exp, ctx.Err().Error())
	}
}
