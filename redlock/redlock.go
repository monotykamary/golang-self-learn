package redlock

import (
	"context"
	"errors"
	"math/rand"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var (
	ErrReleaseResource = errors.New("lock: failed releasing resource")
	ErrExtendResource  = errors.New("lock: failed extending resource")
	ErrAcquireResource = errors.New("lock: failed acquiring resource")
)

var scriptRelease = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
else
	return 0
end
`

var scriptExtend = `
if redis.call("GET", KEYS[1]) ~= ARGV[1] then
  return 0
end

return redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
`

var scriptAcquire = `
if redis.call("EXISTS", KEYS[1]) == 1 then
	return 0
end

return redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
`

type Locker interface {
	Release(ctx context.Context) error
	Extend(ctx context.Context) error
	Acquire(ctx context.Context) error
}

type DLM struct {
	redisClients []*redis.Client
	expiration   time.Duration
	drift        time.Duration
	quorum       int
}

func NewDLM(redisClients []*redis.Client, expiration time.Duration, drift time.Duration) *DLM {
	return &DLM{
		redisClients: redisClients,
		expiration:   expiration,
		drift:        drift,
		quorum:       len(redisClients)/2 + 1,
	}
}

type lock struct {
	redisClients []*redis.Client
	expiration   time.Duration
	drift        time.Duration
	quorum       int
	name         string
	value        string
}

func (dlm *DLM) NewLock(name string) *lock {
	return newLock(dlm, name)
}

func newLock(dlm *DLM, name string) *lock {
	return &lock{
		redisClients: dlm.redisClients,
		quorum:       dlm.quorum,
		name:         name,
		value:        generateRandomString(),
		expiration:   dlm.expiration,
		drift:        dlm.drift,
	}
}

func (l *lock) Release(ctx context.Context) error {
	totalSuccess := 0

	for _, rc := range l.redisClients {
		status, err := rc.Eval(ctx, scriptRelease, []string{l.name}, l.value).Result()

		if err != nil {
			return err
		}

		if status != int64(0) {
			totalSuccess++
		}
	}

	if totalSuccess < l.quorum {
		return ErrReleaseResource
	}

	return nil
}

func (l *lock) Extend(ctx context.Context) error {
	totalSuccess := 0

	for _, rc := range l.redisClients {
		status, err := rc.Eval(ctx, scriptExtend, []string{l.name}, l.value, l.expiration.Microseconds()).Result()

		if err != nil {
			return err
		}

		if status == "OK" {
			totalSuccess++
		}
	}

	if totalSuccess < l.quorum {
		l.Release(ctx)
		return ErrExtendResource
	}

	return nil
}

func (l *lock) Acquire(ctx context.Context) error {
	totalSuccess := 0
	start := time.Now()

	for _, rc := range l.redisClients {
		status, err := rc.Eval(ctx, scriptAcquire, []string{l.name}, l.value, l.expiration.Microseconds()).Result()

		if err != nil {
			return err
		}

		ok := status == "OK"
		now := time.Now()
		isTimeValid := (l.expiration - (now.Sub(start) - l.drift)) > 0

		if ok && isTimeValid {
			totalSuccess++
		}
	}

	if totalSuccess < l.quorum {
		l.Release(ctx)
		return ErrAcquireResource
	}

	return nil
}

func generateRandomString() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, time.Now().Unix()%64)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
