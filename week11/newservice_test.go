package main

import (
	"testing"
	"time"
)

func Test_NewsService_MockBased_Start_Stop(t *testing.T) {
	t.Parallel()

	ns := NewNewService()

	ns.Start()

	ns.Subscribe("Subscriber_1", "Tech")

	ns.Subscribe("Subscriber_2", "Movies")

	ns.Subscribe("Subscriber_3", "Sports")

	ns.MockRegistration("MockBasedSource", "Tech", "Movies", "Sports")

	time.Sleep(time.Second * time.Duration(20))

	ns.Stop()
}

func Test_NewsService_FileBased_Start_Stop(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/newsservice/newsarticles1.json"

	generateFileBasedMultipleCategoryData(filePath)

	ns := NewNewService()

	ns.Start()

	ns.Subscribe("Subscriber_1", "Tech")

	ns.Subscribe("Subscriber_2", "Movies")

	ns.Subscribe("Subscriber_3", "Sports")

	ns.FileBasedRegistration("FileBasedSource", filePath, "Tech", "Movies", "Sports")

	time.Sleep(time.Second * time.Duration(20))

	ns.Stop()
}

func Test_NewsService_MockBased_Registration_UnRegister(t *testing.T) {
	t.Parallel()

	ns := NewNewService()

	ns.Start()

	ns.Subscribe("Subscriber_1", "Tech", "Sports")

	ns.MockRegistration("MockBasedSource_1", "Tech")

	time.Sleep(time.Second * time.Duration(50))

	ns.UnRegister("MockBasedSource_1")

	ns.MockRegistration("MockBasedSource_2", "sports")

	time.Sleep(time.Second * time.Duration(20))

	ns.UnRegister("MockBasedSource_2")

	ns.Stop()
}

func Test_NewsService_FileBased_Registration_UnRegister(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/newsservice/newsarticles2.json"

	generateFileBasedMultipleCategoryData(filePath)

	ns := NewNewService()

	ns.Start()

	ns.Subscribe("Subscriber_2", "Movies")

	ns.FileBasedRegistration("FileBasedSource_1", filePath, "Movies")

	time.Sleep(time.Second * time.Duration(20))

	ns.UnRegister("FileBasedSource_1")

	ns.Stop()
}

func Test_NewsService_Subscribe_Unsubscribe(t *testing.T) {
	t.Parallel()

	ns := NewNewService()

	ns.Start()

	ns.Subscribe("Subscriber_99", "Sports", "Tech")

	ns.UnSubscribe("Subscriber_99")

	ns.Stop()
}

func Test_NewsService_GetArticlesByIds_From_Inmemory(t *testing.T) {
	t.Parallel()

	ns := NewNewService()

	ns.Start()

	ns.Subscribe("Subscriber_1", "Tech")

	ns.Subscribe("Subscriber_2", "Movies")

	ns.Subscribe("Subscriber_3", "Sports")

	ns.MockRegistration("MockBasedSource", "Tech", "Movies", "Sports")

	time.Sleep(time.Second * time.Duration(20))

	var ids []string

	ids = append(ids, "1")
	ids = append(ids, "2")
	ids = append(ids, "3")

	ns.GetArticlesByIds("", ids)

	ns.Stop()
}

func Test_NewsService_GetArticlesByIds_From_BackupFile(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/newsservice/newsarticles3.json"

	generateBackupData(filePath)

	ns := NewNewService()

	var ids []string

	ids = append(ids, "1")
	ids = append(ids, "2")
	ids = append(ids, "3")

	ns.GetArticlesByIds(filePath, ids)

	ns.Stop()
}

func Test_NewsService_GetStreamByCategory_From_Inmemory(t *testing.T) {
	t.Parallel()

	ns := NewNewService()

	ns.Start()

	ns.Subscribe("Subscriber_1", "Tech")

	ns.Subscribe("Subscriber_2", "Movies")

	ns.Subscribe("Subscriber_3", "Sports")

	ns.MockRegistration("MockBasedSource", "Tech", "Movies", "Sports")

	time.Sleep(time.Second * time.Duration(20))

	var categories []string

	categories = append(categories, "Tech")
	categories = append(categories, "Movies")
	categories = append(categories, "Sports")

	ns.GetStreamByCategory("", categories)

	ns.Stop()
}

func Test_NewsService_GetStreamByCategory_From_BackupFile(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/newsservice/newsarticles4.json"

	generateBackupData(filePath)

	ns := NewNewService()

	var categories []string

	categories = append(categories, "Tech")
	categories = append(categories, "Movies")
	categories = append(categories, "Sports")

	ns.GetStreamByCategory(filePath, categories)

	ns.Stop()
}

func Test_NewsService_GetNewsServiceStats(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/newsservice/newsarticles5.json"

	generateBackupData(filePath)

	ns := NewNewService()

	ns.GetNewsServiceStats(filePath)

	ns.Stop()
}

func Test_NewsService_Clear(t *testing.T) {
	t.Parallel()

	filePath := "./testdata/newsservice/newsarticles6.json"

	generateBackupData(filePath)

	ns := NewNewService()

	ns.Clear(filePath)

	ns.Stop()
}
