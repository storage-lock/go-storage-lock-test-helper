package storage_lock_test_helper

import (
	"context"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
	"github.com/storage-lock/go-utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// BasicTest 锁的基本功能测试
func BasicTest[Connection any](t *testing.T, factory *storage_lock_factory.StorageLockFactory[Connection]) {

	lockId := utils.RandomID()
	ownerId := utils.RandomID()

	lock, err := factory.CreateLock(lockId)
	assert.Nil(t, err)

	err = lock.Lock(context.Background(), ownerId)
	assert.Nil(t, err)

	err = lock.UnLock(context.Background(), ownerId)
	assert.Nil(t, err)

	t.Log("StorageLock basic test done")

}
