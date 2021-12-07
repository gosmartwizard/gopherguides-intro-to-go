package week10

import (
	"context"
	"testing"
)

func Test_NewsService_Start_Stop(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ns := NewNewService()

	ns.Start(ctx)

	ns.Stop()
}

func Test_NewsService_Subscribe_Unsubscribe(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ns := NewNewService()

	ns.Start(ctx)

	ns.Subscribe("Subscriber_99", "Sports", "Tech")

	ns.UnSubscribe("Subscriber_99")
}
