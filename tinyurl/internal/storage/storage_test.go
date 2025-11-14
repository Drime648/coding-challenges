package storage

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestStorageInit(t *testing.T) {
	err := InitStorage()
	assert.Nil(t, err)
}

func TestSetAndGet(t *testing.T) {
	original := "www.youtube.com"
	shortened := "abcd"

	err := SaveUrl(original, shortened)
	assert.Nil(t, err)

	resp, err := GetUrl(shortened)
	assert.Nil(t, err)
	assert.Equal(t, original, resp)
}
