package main

import (
	"context"
	"log"
	"time"

	redlock "github.com/monotykamary/golang-self-learn/redlock"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	log.SetFlags(log.Ltime)
	rc1 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7001"})
	rc2 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7002"})
	rc3 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7003"})

	dlm := redlock.NewDLM([]*redis.Client{rc1, rc2, rc3}, 10*time.Second, 2*time.Second)

	// withLockAndUnlock(dlm)
	// withLockAndExtend(dlm)
	withLockOnly(dlm)
}

func withLockAndUnlock(dlm *redlock.DLM) {
	ctx := context.Background()
	locker := dlm.NewLock("this-is-a-key-002")

	if err := locker.Acquire(ctx); err != nil {
		log.Fatal("[main] Failed when locking, err:", err)
	}

	// Perform operation.
	someOperation()

	if err := locker.Release(ctx); err != nil {
		log.Fatal("[main] Failed when unlocking, err:", err)
	}

	log.Println("[main] Done")
}

func withLockAndExtend(dlm *redlock.DLM) {
	ctx := context.Background()
	locker := dlm.NewLock("this-is-a-key-002")

	if err := locker.Acquire(ctx); err != nil {
		log.Fatal("[main] Failed when locking, err:", err)
	}

	// Perform operation.
	someOperation()

	if err := locker.Extend(ctx); err != nil {
		log.Fatal("[main] Failed when extending, err:", err)
	}

	log.Println("[main] Done")
}

func withLockOnly(dlm *redlock.DLM) {
	ctx := context.Background()
	locker := dlm.NewLock("this-is-a-key-002")

	if err := locker.Acquire(ctx); err != nil {
		log.Fatal("[main] Failed when locking, err:", err)
	}

	// Perform operation.
	someOperation()

	// Don't unlock
	log.Println("[main] Done")
}

func someOperation() {
	log.Println("[someOperation] Process has been started")
	time.Sleep(1 * time.Second)
	log.Println("[someOperation] Process has been finished")
}
