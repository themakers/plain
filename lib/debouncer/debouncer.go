package debouncer

import (
	"context"
	"sync"
	"time"
)

type Debouncer struct {
	threshold time.Duration
	limit     time.Duration

	cancel   func()
	flush    func()
	done     context.Context
	makeDone func()
	lock     sync.Mutex
}

func New(threshold, limit time.Duration) *Debouncer {
	tt := &Debouncer{
		threshold: threshold,
		limit:     limit,
		cancel:    func() {},
		flush:     func() {},
	}

	tt.done, tt.makeDone = context.WithCancel(context.Background())
	tt.makeDone()

	return tt
}

func (deb *Debouncer) timer(ctx context.Context, flush context.Context, fn func()) {

	do := func() {
		deb.lock.Lock()
		defer deb.lock.Unlock()
		defer deb.makeDone()

		fn()
	}

	select {
	case <-ctx.Done():
	case <-flush.Done():
		do()
	case <-time.After(deb.threshold):
		do()
	}
}

func (deb *Debouncer) Trigger(fn func()) {
	deb.lock.Lock()
	defer deb.lock.Unlock()

	if deb.cancel != nil {
		deb.cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	deb.cancel = cancel

	ctxFlush, flush := context.WithCancel(context.Background())
	deb.flush = flush

	go deb.timer(ctx, ctxFlush, fn)
}

func (deb *Debouncer) Flush() {
	deb.lock.Lock()

	deb.flush()
    done := deb.done

    deb.lock.Unlock()

	<-done.Done()
}
