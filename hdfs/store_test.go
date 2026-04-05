package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpic"
	pathname := CASPathTransformFun(key)
	assert.Equal(t, "6f90c/0cbff/d1b2a/a1e69/c839a/5b960/6ff14/5c565", pathname)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some data"))
	if err := s.writeStream("test_file", data); err != nil {
		t.Error(err)
	}
}
