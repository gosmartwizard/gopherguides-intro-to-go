package week11

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func generateFileBasedMultipleCategoryData(filePath string) error {

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

func generateFileBasedSingleCategoryData(filePath string) error {

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

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	filePath := "./testdata/filebased/newsarticles1.json"

	generateFileBasedMultipleCategoryData(filePath)

	fileBasedSource := NewFileBasedSource("FileBasedSource", filePath, 10)

	fileBasedSource.SourceStart("Sports", "Tech", "Movies")

	go fileBasedSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-fileBasedSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	cancel()
}

func Test_FileBasedSource_Sports_Category(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	filePath := "./testdata/filebased/newsarticles2.json"

	generateFileBasedSingleCategoryData(filePath)

	fileBasedSource := NewFileBasedSource("FileBasedSource", filePath, 10)

	fileBasedSource.SourceStart("Sports")

	go fileBasedSource.PublishArticles(ctx)

	select {
	case <-ctx.Done():
	case _, ok := <-fileBasedSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	cancel()
}

func Test_FileBasedSource_Start_Stop(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	_, cancel := context.WithCancel(ctx)

	filePath := "./testdata/filebased/newsarticles3.json"

	fileBasedSource := NewFileBasedSource("FileBasedSource", filePath, 10)

	fileBasedSource.SourceStart("Sports", "Tech", "Movies")

	fileBasedSource.SourceStop()

	fileBasedSource.SourceStop()

	cancel()
}

func Test_FileBasedSource_WithTimeOut(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/filebased/newsarticles4.json"

	generateFileBasedMultipleCategoryData(filePath)

	rootCtx := context.Background()

	ctx, cancel := context.WithTimeout(rootCtx, 20*time.Second)

	defer cancel()

	fileBasedSource := NewFileBasedSource("FileBasedSource", filePath, 10)

	fileBasedSource.SourceStart("Sports", "Tech", "Movies")

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
