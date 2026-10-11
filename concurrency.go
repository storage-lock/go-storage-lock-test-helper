package storage_lock_test_helper

import (
	"context"
	"fmt"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
	"github.com/storage-lock/go-utils"
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
)

var (

	// PlayerNum 有多少个人参与
	PlayerNum = 100

	// EveryOnePlayTimes 参与的每个人都操作多少次
	EveryOnePlayTimes = 100
)

// ConcurrencyTest 并发测试
func ConcurrencyTest[Connection any](t *testing.T, factory *storage_lock_factory.StorageLockFactory[Connection]) int {

	lockId := utils.RandomID()
	counter := NewCounter()
	var failures atomic.Int32
	var playerWg sync.WaitGroup
	for i := 0; i < PlayerNum; i++ {
		playerWg.Add(1)
		go func() {
			defer playerWg.Done()

			player, err := createPlayer(lockId, factory)
			if err != nil {
				failures.Add(1)
				return
			}

			for j := 0; j < EveryOnePlayTimes; j++ {

				err = runPlay(player, counter)

				if err != nil {
					failures.Add(1)
					return
				}

			}

		}()
	}
	playerWg.Wait()
	failed := int(failures.Load())
	counter.Close()
	counter.Wait()
	if failed == 0 {
		c := counter.Get()
		assert.Equal(t, PlayerNum*EveryOnePlayTimes, c)
	}
	t.Log("StorageLock concurrency test done")
	return failed
}

var createPlayer = func(lockId string, factory any) (*LockPlayer[any], error) {
	return NewLockPlayer(lockId, factory.(*storage_lock_factory.StorageLockFactory[any]))
}

var runPlay = func(player *LockPlayer[any], counter *Counter) error {
	return player.Do(context.Background(), func() {
		fmt.Println(counter.Get())
		counter.Add(1)
	})
}
