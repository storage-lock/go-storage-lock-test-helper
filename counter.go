package storage_lock_test_helper

import "sync"

// Counter 计数器
type Counter struct {
	count   int
	channel chan int
	wg      sync.WaitGroup
}

func NewCounter() *Counter {
	c := &Counter{
		count:   0,
		channel: make(chan int, 100),
		wg:      sync.WaitGroup{},
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for n := range c.channel {
			c.count += n
		}
	}()

	return c
}

func (x *Counter) Add(n int) {
	x.channel <- n
}

func (x *Counter) Get() int {
	return x.count
}

func (x *Counter) Wait() {
	x.wg.Wait()
}

func (x *Counter) Close() {
	close(x.channel)
}
