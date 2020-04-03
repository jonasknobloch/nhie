package cache

import (
	"bytes"
	"encoding/gob"
)

type gobSerializer struct{}

func newGobSerializer() *gobSerializer {
	return &gobSerializer{}
}

func (g *gobSerializer) serialize(c interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(c)

	if err != nil {
		return nil, newSerializeError(err)
	}

	return buf.Bytes(), nil
}

func (g *gobSerializer) deserialize(d []byte, c interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(d))
	err := dec.Decode(c)

	if err != nil {
		return newSerializeError(err)
	}

	return nil
}
