package cache

import "C"
import (
	"errors"
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

func Retrieve(k Key) (interface{}, error) {
	g := C.Get(k.String())

	if g.Val() == "" {
		return nil, errors.New("not in cache")
	}

	var i interface{}

	serializer := newGobSerializer()
	err := serializer.deserialize([]byte(g.String()), &i)

	if err != nil {
		return nil, err
	}

	return i, nil
}

func Store(k Key, i interface{}, exp time.Duration) error {
	serializer := newGobSerializer()
	b, err := serializer.serialize(i)

	if err != nil {
		return err
	}

	C.Set(k.String(), b, exp)
	return nil
}

func Clear(k Key) {
	C.Del(C.Keys(k.String()).Val()...)
}
