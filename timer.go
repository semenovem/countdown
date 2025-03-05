package countdown

import (
	"context"
	"time"
)

type Timer struct {
	c    chan struct{}
	ctx  context.Context
	d    time.Duration
	stop bool
}

func NewTimer(ctx context.Context, d time.Duration) *Timer {
	t := &Timer{
		c:   make(chan struct{}),
		d:   d,
		ctx: ctx,
	}

	go func() {
		<-ctx.Done()
		t.stop = true
		t.c = nil
	}()

	return t
}

func (t *Timer) Start() {
	if t.c == nil {
		panic("timer not initialized")
	}

	go func() {
		time.Sleep(t.d)
		if t.ctx.Err() == nil && !t.stop {
			t.c <- struct{}{}
		}
	}()
}

func (t *Timer) Stop() {
	if t.c == nil {
		panic("timer not initialized")
	}
	t.stop = true
}

func (t *Timer) C() <-chan struct{} {
	if t.c == nil {
		t.c = make(chan struct{})
	}

	return t.c
}
