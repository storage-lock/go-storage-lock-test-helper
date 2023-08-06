package storage_lock_test_helper

import (
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
	"testing"
)

// TestStorageLock 测试锁
func TestStorageLock[Connection any](t *testing.T, factory *storage_lock_factory.StorageLockFactory[Connection]) {
	// TODO 2023-8-7 00:47:43 编写锁的测试用例
	factory.CreateLock()
}
