package cache

import (
	"github.com/go-redis/redis"
	"strings"
	"time"
)

type Key []string

func (k Key) String() string {
	if len(k) == 0 {
		panic("key can't be empty")
	}

	return strings.Join(k[:], ":")
}

func Exists(k Key) (bool, error) {
	if err := C.Get(k.String()).Err(); err != nil {
		if err != redis.Nil {
			return false, newError(err)
		}

		return false, nil
	}

	return true, nil
}

func Retrieve(k Key) (string, error) {
	g := C.Get(k.String())

	if err := g.Err(); err != nil {
		if err == redis.Nil {
			return "", newError(ErrKeyNotFound)
		}

		return "", newError(err)
	}

	return g.Val(), nil
}

func Store(k Key, i string, exp time.Duration) error {
	s := C.Set(k.String(), i, exp)

	return s.Err()
}

func Clear(k Key) {
	C.Del(C.Keys(k.String()).Val()...)
}
