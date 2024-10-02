package log

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFrame(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	go frame(ctx, &wg, "waiting on foo", runner)

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}

func TestWait(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := wait(ctx, "waiting on foo", runner)
	defer wg.Wait()

	fmt.Println("waiting on foo")
	time.Sleep(2 * time.Second)
}
