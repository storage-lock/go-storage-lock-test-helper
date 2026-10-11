package storage_lock_test_helper

import (
	"context"
	"errors"
	"testing"

	memory_locks "github.com/storage-lock/go-memory-locks"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

func TestHelperWithMemoryFactory(t *testing.T) {
	oldPlayers, oldTimes := PlayerNum, EveryOnePlayTimes
	PlayerNum, EveryOnePlayTimes = 2, 2
	t.Cleanup(func() {
		PlayerNum, EveryOnePlayTimes = oldPlayers, oldTimes
	})
	factory, err := memory_locks.NewLockFactory()
	if err != nil {
		t.Fatal(err)
	}
	if err := TestStorageLock(t, factory.StorageLockFactory); err != nil {
		t.Fatal(err)
	}

	counter := NewCounter()
	counter.Add(3)
	counter.Close()
	counter.Wait()
	if counter.Get() != 3 {
		t.Fatal(counter.Get())
	}

	if _, err := NewLockPlayer[any]("bad", &storage_lock_factory.StorageLockFactory[any]{}); err == nil {
		t.Fatal("nil storage factory accepted")
	}
	bad, err := factory.CreateLock("player-error")
	if err != nil {
		t.Fatal(err)
	}
	player := &LockPlayer[any]{ownerId: "owner", lock: bad}
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	if err := player.Do(canceled, func() { t.Fatal("callback ran") }); err == nil {
		t.Fatal("canceled lock succeeded")
	}
}

func TestConcurrencyFailureBranches(t *testing.T) {
	oldPlayers, oldTimes := PlayerNum, EveryOnePlayTimes
	oldCreate, oldPlay := createPlayer, runPlay
	PlayerNum, EveryOnePlayTimes = 1, 1
	t.Cleanup(func() {
		PlayerNum, EveryOnePlayTimes = oldPlayers, oldTimes
		createPlayer, runPlay = oldCreate, oldPlay
	})
	factory, err := memory_locks.NewLockFactory()
	if err != nil {
		t.Fatal(err)
	}
	createPlayer = func(string, any) (*LockPlayer[any], error) { return nil, errors.New("create failed") }
	if err := TestStorageLock(t, factory.StorageLockFactory); err == nil {
		t.Fatal("create failure was accepted")
	}
	createPlayer = oldCreate
	runPlay = func(*LockPlayer[any], *Counter) error { return errors.New("play failed") }
	if ConcurrencyTest(t, factory.StorageLockFactory) == 0 {
		t.Fatal("play failure ignored")
	}
}
