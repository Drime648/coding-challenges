package main

import (
	"bytes"
	"testing"
)

// func TestPathTransformFunc(t *testing.T) {

// }

func TestStore(t *testing.T) {
	opts := StoreOpts {
		PathTransformFunc: DefaultPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some data"))
	if err := s.writeStream("test_file", data); err != nil {
		t.Error(err)
	}
}
