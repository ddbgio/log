package log

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestFrame(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go frame(ctx, &wg, "waiting on foo", runner)

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}
