package week10

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"sync"
)

type Configuration struct {
	SaveStateInterval  int `json:"saveStateInterval"`
	MockSourceInterval int `json:"mockSourceInterval"`
	FileBasedInterval  int `json:"FileBasedInterval"`
}

type NewsService struct {
	sync.RWMutex
	ctx                 context.Context
	sources             map[string]interface{}
	subscribers         map[string]*Subscriber
	channelSubscriber   map[string]chan Article
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

func NewNewService() *NewsService {
	ns := &NewsService{}

	ns.closed = false

	ns.subscribers = make(map[string]*Subscriber)
	ns.channelSubscriber = make(map[string]chan Article)

	ns.sources = make(map[string]interface{})

	ns.categorySubscribers = make(map[string][]string)

	ns.ctx = context.Background()

	ns.newsArticles = &NewsArticles{}
	ns.newsArticles.newsArticles = make(map[int]Article)

	ns.categoryArticles = &CategoryArticles{}
	ns.categoryArticles.categoryArticles = make(map[string]map[int]Article)

	ns.newsStats = &NewsStats{}
	ns.newsStats.backupFileLocation = "./newServiceBackupFile.json"
	ns.newsStats.articlesPerCategory = make(map[string]int)
	ns.newsStats.articlesPerSource = make(map[string]int)

	return ns
}

func (ns *NewsService) Start(ctx context.Context) {

	ctx, cancel := context.WithCancel(ctx)

	ns.cancel = cancel

	ns.ctx = ctx

	ns.readConfiguration()

	ns.LoadArticlesFromBackupFile()
}

func (ns *NewsService) listenForArticles(news chan []Article) {

	for {
		select {
		case <-ns.ctx.Done():
			ns.Stop()
		case articles, ok := <-news:
			if !ok {
				continue
			}

			for _, article := range articles {
				ns.publish(article)

				ns.saveArticleInMemory(article)
			}
		}
	}
}

func (ns *NewsService) publish(article Article) {

	ns.Lock()
	defer ns.Unlock()

	topic := article.Category

	ss := ns.categorySubscribers[topic]

	for _, s := range ss {
		ns.channelSubscriber[s] <- article
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

func (ns *NewsService) Register(name string, categories ...string) {

	if name == "MockSource" {
		mocksource := NewMockSource(name)

		mocksource.SourceStart(ns.config.MockSourceInterval, categories...)

		ns.Lock()
		ns.sources[name] = mocksource
		ns.Unlock()

		go ns.listenForArticles(mocksource.News)

		go mocksource.PublishArticles()

	} else if name == "FileBasedSource" {
		fileBasedSource := NewFileBasedSource(name)

		fileBasedSource.SourceStart(ns.config.FileBasedInterval, categories...)

		ns.Lock()
		ns.sources[name] = fileBasedSource
		ns.Unlock()

		go ns.listenForArticles(fileBasedSource.News)

		go fileBasedSource.PublishArticles()
	}

}
func (ns *NewsService) UnRegister(name string) {

	ns.Lock()
	defer ns.Unlock()

	source := ns.sources[name]

	if s, ok := source.(MockSource); ok {
		s.SourceStop()
	} else if s, ok := source.(FileBasedSource); ok {
		s.SourceStop()
	}

	delete(ns.sources, name)
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
	news.channelSubscriber[newSubscriber.Name] = ch

	news.Unlock()

	newSubscriber.Listen(ch)
}

func (ns *NewsService) UnSubscribe(name string) {

	ns.Lock()
	defer ns.Unlock()

	subscriber, ok := ns.subscribers[name]

	if ok {
		ch := ns.channelSubscriber[subscriber.Name]
		if ch != nil {
			close(ch)
		}
		delete(ns.channelSubscriber, name)

		subscriber.Cancel()
		delete(ns.subscribers, name)
	}
}

func (ns *NewsService) saveArticlesInBackupFile() error {

	ns.RLock()
	defer ns.RUnlock()

	fileBytes, err := json.Marshal(ns.newsArticles.newsArticles)

	if err != nil {
		return err
	}

	ioutil.WriteFile(ns.newsStats.backupFileLocation, fileBytes, 0644)

	if err != nil {
		return err
	}

	return nil
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

	fileBytes, err := ioutil.ReadFile("./config/configuration.json")

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

		ns.cancel()

		ns.closed = true

		for _, source := range ns.sources {
			if s, ok := source.(MockSource); ok {
				s.SourceStop()
			} else if s, ok := source.(FileBasedSource); ok {
				s.SourceStop()
			}
		}

		for name, ch := range ns.channelSubscriber {
			if ch != nil {
				close(ch)
			}
			delete(ns.channelSubscriber, name)
		}

		for name, s := range ns.subscribers {
			s.Cancel()
			delete(ns.subscribers, name)
		}
	})
}

/* func (ns *NewsService) NewsServiceStats() {

	ns.RLock()
	defer ns.RUnlock()

	fmt.Println("NewsArticles")
	for index, article := range ns.newsArticles.newsArticles {
		fmt.Printf("Index : %#v, Article : %#v \n", index, article)
	}

	fmt.Println()

	fmt.Println("CategoryArticles")
	for name, category := range ns.categoryArticles.categoryArticles {
		for index, article := range category {
			fmt.Printf("Category : %#v, Index : %#v, Article : %#v \n", name, index, article)
		}
	}

	fmt.Println("NewsStats")
	fmt.Printf("ArticlesPerCategory : %#v \n", ns.newsStats.articlesPerCategory)
	fmt.Printf("ArticlesPerSource : %#v \n", ns.newsStats.articlesPerSource)
	fmt.Printf("TotalArticles : %#v \n", ns.newsStats.totalArticles)
	fmt.Printf("Categories : %#v \n", ns.newsStats.categories)
	fmt.Println()
} */
