package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpic"
	pathname := CASPathTransformFun(key)
	expectedPathName := "6f90c/0cbff/d1b2a/a1e69/c839a/5b960/6ff14/5c565"
	assert.Equal(t, expectedPathName, pathname)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFun,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some data"))
	assert.Nil(t, s.writeStream("test_file", data))
}
