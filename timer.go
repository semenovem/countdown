package countdown

import (
	"time"
)

type Timer struct {
	c   chan struct{}
	d   time.Duration
	ind int
}

func NewTimer(d time.Duration) *Timer {
	t := &Timer{
		c: make(chan struct{}),
		d: d,
	}

	return t
}

func (t *Timer) Start() {
	if t.c == nil {
		panic("timer not initialized")
	}

	i := t.ind

	go func() {
		<-time.After(t.d)

		if i == t.ind {
			t.c <- struct{}{}
		}
	}()
}

func (t *Timer) Stop() {
	if t.c == nil {
		panic("timer not initialized")
	}

	t.ind++
}

func (t *Timer) C() <-chan struct{} {
	if t.c == nil {
		panic("timer not initialized")
	}

	return t.c
}
