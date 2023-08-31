package storage_lock_test_helper

import (
	"context"
	"fmt"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var (

	// PlayerNum 有多少个人参与
	PlayerNum = 100

	// EveryOnePlayTimes 参与的每个人都操作多少次
	EveryOnePlayTimes = 100
)

// ConcurrencyTest 并发测试
func ConcurrencyTest[Connection any](t *testing.T, factory *storage_lock_factory.StorageLockFactory[Connection]) {

	lockId := "test-lock"

	counter := NewCounter()
	var playerWg sync.WaitGroup
	for i := 0; i < PlayerNum; i++ {
		playerWg.Add(1)
		go func() {
			defer playerWg.Done()

			for i := 0; i < EveryOnePlayTimes; i++ {

				player, err := NewLockPlayer[Connection](lockId, factory)
				if err != nil {
					panic(err)
				}

				err = player.Do(context.Background(), func() {
					fmt.Println(counter.Get())
					counter.Add(1)
				})

				if err != nil {
					panic(err)
				}

			}

		}()
	}

	playerWg.Wait()
	counter.Close()
	counter.Wait()
	c := counter.Get()
	assert.Equal(t, PlayerNum*EveryOnePlayTimes, c)

}
