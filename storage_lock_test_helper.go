package storage_lock_test_helper

import (
	"fmt"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
	"testing"
)

// TestStorageLock 测试锁
func TestStorageLock[Connection any](t *testing.T, factory *storage_lock_factory.StorageLockFactory[Connection]) error {

	// 基础功能测试
	BasicTest(t, factory)

	// 并发测试
	if failed := ConcurrencyTest(t, factory); failed != 0 {
		return fmt.Errorf("players failed: %d", failed)
	}
	t.Log("StorageLock all testcase done.")
	return nil
}
