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

	generateMultipleCategoryData("./newsarticles/newsarticles.json")

	fileBasedSource := NewFileBasedSource("FileBasedSource")

	fileBasedSource.SourceStart(10, "Sports")

	go fileBasedSource.PublishArticles()

	select {
	case <-fileBasedSource.ctx.Done():
	case _, ok := <-fileBasedSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	fileBasedSource.Cancel()
}

func Test_FileBasedSource_Sports_Category(t *testing.T) {
	t.Parallel()

	generateSingleCategoryData("./newsarticles/newsarticles.json")

	fileBasedSource := NewFileBasedSource("FileBasedSource")

	fileBasedSource.SourceStart(10, "Tech")

	go fileBasedSource.PublishArticles()

	select {
	case <-fileBasedSource.ctx.Done():
	case _, ok := <-fileBasedSource.News:
		if !ok {
			t.Fatalf("Expected : Open Channel, got : closed Channel")
		}
	}

	fileBasedSource.Cancel()
}

func Test_FileBasedSource_Start_Stop(t *testing.T) {
	t.Parallel()

	fileBasedSource := NewFileBasedSource("FileBasedSource")

	fileBasedSource.SourceStart(10, "Sports", "Tech", "Movies")

	fileBasedSource.SourceStop()

	fileBasedSource.SourceStop()
}

func Test_FileBasedSource_WithTimeOut(t *testing.T) {
	t.Parallel()

	generateMultipleCategoryData("./testdata/newsarticles.json")

	rootCtx := context.Background()

	ctx, cancel := context.WithTimeout(rootCtx, 30*time.Second)

	defer cancel()

	fileBasedSource := NewFileBasedSource("FileBasedSource")

	fileBasedSource.SourceStart(10, "Sports", "Tech", "Movies")

	go fileBasedSource.PublishArticles()

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
