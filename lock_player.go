package storage_lock_test_helper

import (
	"context"
	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
	"github.com/storage-lock/go-utils"
)

// LockPlayer 竞争锁的人
type LockPlayer[Connection any] struct {
	lockId  string
	ownerId string
	lock    *storage_lock.StorageLock
	factory *storage_lock_factory.StorageLockFactory[Connection]
}

func NewLockPlayer[Connection any](lockId string, factory *storage_lock_factory.StorageLockFactory[Connection]) (*LockPlayer[Connection], error) {
	lock, err := factory.CreateLock(lockId)
	if err != nil {
		return nil, err
	}
	return &LockPlayer[Connection]{
		lockId:  lockId,
		ownerId: utils.RandomID(),
		lock:    lock,
		factory: factory,
	}, nil
}

func (x *LockPlayer[Connection]) Do(ctx context.Context, f func()) error {
	err := x.lock.Lock(ctx, x.ownerId)
	if err != nil {
		return err
	}
	f()
	return x.lock.UnLock(ctx, x.ownerId)
}
