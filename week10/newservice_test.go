package week10

/* func Test_NewsService_1(t *testing.T) {
	t.Parallel()

	ns := NewNewService()

	ns.Start()

	ss := ns.categorySubscribers["Sports"]

	for _, s := range ss {
		sub := ns.subscribers[s]
		sub.GetChannel() <- "Hello Sports from NewsService"
	}
} */

/* func Test_NewsService_2(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ns := NewNewService()

	go ns.Start(ctx)

	time.Sleep(time.Millisecond * 5000)

	//ns.NewsServiceStats()
	//ns.saveArticlesInBackupFile()

	//ns.LoadArticlesFromBackupFile()

	<-ns.ctx.Done()
}
*/
