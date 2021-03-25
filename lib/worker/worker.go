package worker

import "context"

type Worker struct {
	queue chan func()
}

func New() *Worker {
	wor := &Worker{
		queue: make(chan func()),
	}

	return wor
}

func (wor *Worker) Work(ctx context.Context, done func()) {
	defer done()

	for {
		select {
		case <-ctx.Done():
			return
		case task := <-wor.queue:
			task()
		}
	}
}

func (wor *Worker) Run(task func()) {
	wor.queue <- task
}

func (wor *Worker) RunSync(task func()) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wor.queue <- func() {
		defer cancel()
		task()
	}

	<-ctx.Done()
}
