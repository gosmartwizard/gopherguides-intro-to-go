package week11

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
)

type Configuration struct {
	SaveStateInterval  int `json:"saveStateInterval"`
	MockSourceInterval int `json:"mockSourceInterval"`
	FileBasedInterval  int `json:"fileBasedInterval"`
}

type NewsService struct {
	sync.RWMutex
	ctx                 context.Context
	sources             map[string]context.CancelFunc
	subscribers         map[string]*Subscriber
	subscriberChannel   map[string]chan Article
	categorySubscribers map[string][]string
	newsArticles        *NewsArticles
	categoryArticles    *CategoryArticles
	newsStats           *NewsStats
	cancel              context.CancelFunc
	closed              bool
	stopOnce            sync.Once
	config              Configuration
}

type NewsArticles struct {
	newsArticles map[int]Article
}

type CategoryArticles struct {
	categoryArticles map[string]map[int]Article
}

type NewsStats struct {
	categories          []string
	backupFileLocation  string
	totalArticles       int
	articlesPerCategory map[string]int
	articlesPerSource   map[string]int
}

const configFilePath = "./config/configuration.json"

func NewNewService() *NewsService {
	ns := &NewsService{}

	ns.closed = false

	ns.sources = make(map[string]context.CancelFunc)

	ns.subscribers = make(map[string]*Subscriber)
	ns.subscriberChannel = make(map[string]chan Article)

	ns.categorySubscribers = make(map[string][]string)

	ns.ctx = context.Background()

	ns.newsArticles = &NewsArticles{}
	ns.newsArticles.newsArticles = make(map[int]Article)

	ns.categoryArticles = &CategoryArticles{}
	ns.categoryArticles.categoryArticles = make(map[string]map[int]Article)

	ns.newsStats = &NewsStats{}
	ns.newsStats.totalArticles = 0
	ns.newsStats.backupFileLocation = "/tmp/NewsServiceBackup.json"
	ns.newsStats.articlesPerCategory = make(map[string]int)
	ns.newsStats.articlesPerSource = make(map[string]int)

	return ns
}

func (ns *NewsService) Start() error {

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	ns.cancel = cancel

	ns.ctx = ctx

	err := ns.readConfiguration()

	if err != nil {
		return err
	}

	ns.LoadArticlesFromBackupFile()

	go ns.SaveArtilces()

	return nil
}

func (ns *NewsService) listenForArticles(news chan []Article) {

	for {
		select {
		case <-ns.ctx.Done():
			return
		case articles, ok := <-news:
			if !ok {
				return
			}

			for _, article := range articles {
				go ns.publish(article)

				ns.saveArticleInMemory(article)
			}
		}
	}
}

func (ns *NewsService) publish(article Article) {

	ns.RLock()
	defer ns.RUnlock()

	topic := article.Category

	ss := ns.categorySubscribers[topic]

	for _, s := range ss {
		ns.subscriberChannel[s] <- article
	}
}

func (ns *NewsService) saveArticleInMemory(article Article) {

	ns.Lock()
	defer ns.Unlock()

	ns.newsStats.totalArticles += 1

	count := ns.newsStats.totalArticles

	ns.newsArticles.newsArticles[count] = article

	c, ok := ns.categoryArticles.categoryArticles[article.Category]

	if !ok {
		c = make(map[int]Article)
		c[count] = article
		ns.categoryArticles.categoryArticles[article.Category] = c
		ns.newsStats.categories = append(ns.newsStats.categories, article.Category)
	} else {
		c[count] = article
	}

	ns.newsStats.articlesPerCategory[article.Category] += 1
	ns.newsStats.articlesPerSource[article.Source] += 1
}

func (ns *NewsService) MockRegistration(name string, categories ...string) {

	mocksource := NewMockSource(name, ns.config.MockSourceInterval)

	mocksource.SourceStart(categories...)

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	ns.Lock()
	ns.sources[name] = cancel
	ns.Unlock()

	go ns.listenForArticles(mocksource.News)

	go mocksource.PublishArticles(ctx)

}

func (ns *NewsService) FileBasedRegistration(name string, filePath string, categories ...string) {

	fileBasedSource := NewFileBasedSource(name, filePath, ns.config.FileBasedInterval)

	ch := fileBasedSource.SourceStart(categories...)

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	ns.Lock()
	ns.sources[name] = cancel
	ns.Unlock()

	go ns.listenForArticles(ch)

	go fileBasedSource.PublishArticles(ctx)

}

func (ns *NewsService) UnRegister(name string) {

	ns.Lock()
	defer ns.Unlock()

	cancel, ok := ns.sources[name]
	if ok {
		cancel()
		delete(ns.sources, name)
	}
}

func (news *NewsService) Subscribe(name string, topics ...string) {

	newSubscriber := NewSubscriber(name)

	newSubscriber.SubscriberStart(topics...)

	news.Lock()

	news.subscribers[name] = newSubscriber

	for _, topic := range topics {
		news.categorySubscribers[topic] = append(news.categorySubscribers[topic], newSubscriber.Name)
	}

	ch := make(chan Article)
	news.subscriberChannel[newSubscriber.Name] = ch

	news.Unlock()

	newSubscriber.Listen(ch)
}

func (ns *NewsService) UnSubscribe(name string) {

	ns.Lock()
	defer ns.Unlock()

	subscriber, ok := ns.subscribers[name]

	if ok {
		ch := ns.subscriberChannel[subscriber.Name]
		if ch != nil {
			close(ch)
		}
		delete(ns.subscriberChannel, name)

		subscriber.Cancel()
		delete(ns.subscribers, name)
	}
}

func (ns *NewsService) SaveArtilces() {
	for {
		time.Sleep(time.Second * time.Duration(ns.config.SaveStateInterval))
		ns.saveArticlesInBackupFile()
	}
}

func (ns *NewsService) saveArticlesInBackupFile() error {

	ns.RLock()
	defer ns.RUnlock()

	if len(ns.newsArticles.newsArticles) == 0 {
		return nil
	}

	bytes, err := json.Marshal(ns.newsArticles.newsArticles)

	if err != nil {
		return err
	}

	file, err := os.Create(ns.newsStats.backupFileLocation)

	if err != nil {
		return err
	}

	_, err = file.Write(bytes)

	return err
}

func (ns *NewsService) LoadArticlesFromBackupFile() error {

	fileBytes, err := ioutil.ReadFile(ns.newsStats.backupFileLocation)

	if err != nil {
		return err
	}

	var articles map[int]Article

	err = json.Unmarshal(fileBytes, &articles)

	if err != nil {
		return err
	}

	for _, article := range articles {
		ns.saveArticleInMemory(article)
	}

	return nil
}

func (ns *NewsService) readConfiguration() error {

	_, err := os.Open(configFilePath)
	if errors.Is(err, os.ErrNotExist) {
		return errors.New("config file doesn't exists")
	}

	fileBytes, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		return err
	}

	var config Configuration

	err = json.Unmarshal(fileBytes, &config)

	if err != nil {
		return err
	}

	ns.config = config

	return nil
}

func (ns *NewsService) Stop() {

	ns.RLock()
	if ns.closed {
		ns.RUnlock()
		return
	}
	ns.RUnlock()

	ns.stopOnce.Do(func() {

		ns.saveArticlesInBackupFile()

		ns.Lock()
		defer ns.Unlock()

		if ns.cancel != nil {
			ns.cancel()
		}

		ns.closed = true

		for name, cancel := range ns.sources {
			cancel()
			delete(ns.sources, name)
		}

		for name, ch := range ns.subscriberChannel {
			if ch != nil {
				close(ch)
			}
			delete(ns.subscriberChannel, name)
		}

		for name, s := range ns.subscribers {
			s.Cancel()
			delete(ns.subscribers, name)
		}
	})
}

func (ns *NewsService) GetArticlesByIds(backupFile string, articleIds []string) (map[int]Article, error) {

	articles := make(map[int]Article)

	if ns != nil {
		ns.RLock()
		for _, id := range articleIds {

			id, err := strconv.Atoi(id)

			if err != nil {
				return articles, err
			}

			article, ok := ns.newsArticles.newsArticles[id]

			if !ok {
				continue
			}

			articles[id] = article
		}
		ns.RUnlock()

		if len(articleIds) == len(articles) {
			return articles, nil
		}
	}

	_, err := os.Open(backupFile)
	if errors.Is(err, os.ErrNotExist) {
		return articles, errors.New("BackupFile not exists")
	}

	fileBytes, err := ioutil.ReadFile(backupFile)

	if err != nil {
		return articles, err
	}

	var idArticles map[int]Article

	err = json.Unmarshal(fileBytes, &idArticles)

	if err != nil {
		return articles, err
	}

	for _, id := range articleIds {

		id, _ := strconv.Atoi(id)

		article, ok := idArticles[id]

		if !ok {
			return articles, fmt.Errorf("article id : %#v doesn't exist", id)
		}

		articles[id] = article
	}

	return articles, nil
}

func (ns *NewsService) GetStreamByCategory(backupFile string, categories []string) (map[int]Article, error) {

	articles := make(map[int]Article)
	ns.RLock()
	for _, category := range categories {
		newsArticles, ok := ns.categoryArticles.categoryArticles[category]
		if ok {
			for id, article := range newsArticles {
				articles[id] = article
			}
		}
	}
	ns.RUnlock()

	if len(articles) != 0 {
		return articles, nil
	}

	_, err := os.Open(backupFile)
	if errors.Is(err, os.ErrNotExist) {
		return articles, errors.New("BackupFile not exists")
	}

	fileBytes, err := ioutil.ReadFile(backupFile)

	if err != nil {
		return articles, err
	}

	var idArticles map[int]Article

	err = json.Unmarshal(fileBytes, &idArticles)

	if err != nil {
		return articles, err
	}

	for id, article := range idArticles {

		for _, category := range categories {
			if article.Category == category {
				articles[id] = article
				break
			}
		}
	}

	return articles, nil
}

func (ns *NewsService) GetNewsServiceStats(backupFile string) (string, error) {

	_, err := os.Open(backupFile)
	if errors.Is(err, os.ErrNotExist) {
		return "", errors.New("BackupFile not exists")
	}

	bb := &bytes.Buffer{}

	fileBytes, err := ioutil.ReadFile(backupFile)

	if err != nil {
		return bb.String(), err
	}

	var idArticles map[int]Article

	err = json.Unmarshal(fileBytes, &idArticles)

	if err != nil {
		return bb.String(), err
	}

	articlesPerCategory := make(map[string]int)
	articlesPerSource := make(map[string]int)

	articlesCount := len(idArticles)

	for _, article := range idArticles {
		articlesPerCategory[article.Category] += 1
		articlesPerSource[article.Source] += 1
	}

	fmt.Fprintf(bb, "List of categories in the Backup file are as follows\n")
	for category := range articlesPerCategory {
		fmt.Fprintf(bb, "\t %v ", category)
	}
	fmt.Fprintln(bb)

	fmt.Fprintln(bb)
	fmt.Fprintf(bb, "Location of the backup file : %v\n", backupFile)

	fmt.Fprintln(bb)
	fmt.Fprintf(bb, "Number of articles in the backup file : %v\n", articlesCount)

	fmt.Fprintln(bb)
	fmt.Fprintf(bb, "Number of articles per Category are as follows \n")
	for category, count := range articlesPerCategory {
		fmt.Fprintf(bb, " %v : %v  \n", category, count)
	}

	fmt.Fprintln(bb)

	fmt.Fprintf(bb, "Number of articles per Source are as follows\n")
	for source, count := range articlesPerSource {
		fmt.Fprintf(bb, " %v : %v  \n", source, count)
	}

	fmt.Fprintln(bb)

	return bb.String(), nil
}

func (ns *NewsService) Clear(backupFile string) error {

	_, err := os.Open(backupFile)
	if errors.Is(err, os.ErrNotExist) {
		return errors.New("BackupFile not exists")
	}

	err = os.Remove(backupFile)

	if err != nil {
		return err
	}

	return nil
}
