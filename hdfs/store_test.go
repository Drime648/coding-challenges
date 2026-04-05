package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpic"
	pathKey := CASPathTransformFun(key)
	expectedOriginal := "6f90c0cbffd1b2aa1e69c839a5b9606ff145c565"
	expectedPathName := "6f90c/0cbff/d1b2a/a1e69/c839a/5b960/6ff14/5c565"
	expectedPathKey := PathKey{
		Pathname: expectedPathName,
		Filename: expectedOriginal,
	}
	assert.Equal(t, expectedPathKey, pathKey)
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFun,
	}
	s := NewStore(opts)
	key := "special_key"

	data := []byte("some data")
	assert.Nil(t, s.writeStream(key, bytes.NewReader(data)))

	assert.Nil(t, s.Delete(key))

	_, err := s.Read(key)
	assert.NotNil(t, err)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFun,
	}
	s := NewStore(opts)
	key := "special_key"

	data := []byte("some data")
	assert.Nil(t, s.writeStream(key, bytes.NewReader(data)))

	r, err := s.Read(key)
	assert.Nil(t, err)

	b, err := io.ReadAll(r)

	assert.Equal(t, data, b)
}

