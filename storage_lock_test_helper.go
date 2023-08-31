package storage_lock_test_helper

import (
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
	"testing"
)

// TestStorageLock 测试锁
func TestStorageLock[Connection any](t *testing.T, factory *storage_lock_factory.StorageLockFactory[Connection]) {

	// 基础功能测试
	BasicTest(t, factory)

	// 并发测试
	ConcurrencyTest(t, factory)

}
