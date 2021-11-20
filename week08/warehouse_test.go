package week08

import (
	"context"
	"testing"
)

func Test_Warehouse_Stop(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	w := Warehouse{}

	w.Start(ctx)

	w.Stop()
}
