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

func Retrieve(k Key, i interface{}) error {
	g := C.Get(k.String())

	if err := g.Err(); err != nil {
		if err == redis.Nil {
			return newError(ErrKeyNotFound)
		}

		return newError(err)
	}

	serializer := newGobSerializer()
	err := serializer.deserialize([]byte(g.Val()), i)

	if err != nil {
		return newError(err)
	}

	return nil
}

func Store(k Key, i interface{}, exp time.Duration) error {
	serializer := newGobSerializer()
	b, err := serializer.serialize(i)

	if err != nil {
		return newError(err)
	}

	s := C.Set(k.String(), b, exp)
	return s.Err()
}

func Clear(k Key) {
	C.Del(C.Keys(k.String()).Val()...)
}
