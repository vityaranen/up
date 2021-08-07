package up

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"
)

var (
	ctx context.Context
	wg  *sync.WaitGroup
)

func init() {
	ctx, _ = signal.NotifyContext(context.Background(), os.Interrupt)
	wg = &sync.WaitGroup{}
}

func Ctx() context.Context {
	return ctx
}

func Add(delta int) {
	wg.Add(delta)
}

func Done() {
	wg.Done()
}

func WaitDone(timeout time.Duration) (byTimeout bool) {
	<-ctx.Done()

	ok := make(chan struct{})
	go func() {
		wg.Wait()
		ok <- struct{}{}
	}()

	select {
	case <-ok:
		return false
	case <-time.After(timeout):
		return true
	}
}
