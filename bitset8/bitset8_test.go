package bitset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	b := New()

	ok, err := b.Contains(1)
	assert.Equal(t, nil, err)
	assert.False(t, ok)

	b.Add(1)

	ok, err = b.Contains(1)
	assert.Equal(t, nil, err)
	assert.True(t, ok)
}

func TestContains(t *testing.T) {
	b := New()

	ok, err := b.Contains(10)
	assert.Equal(t, ErrNumberNotFit8, err)
	assert.False(t, ok)

	ok, err = b.Contains(1)
	assert.Equal(t, nil, err)
	assert.False(t, ok)

	b.Add(1)

	ok, err = b.Contains(1)
	assert.Equal(t, nil, err)
	assert.True(t, ok)
}

func TestDelete(t *testing.T) {
	b := New()

	err := b.Delete(10)
	assert.Equal(t, ErrNumberNotFit8, err)

	b.Add(1)

	ok, err := b.Contains(1)
	assert.Equal(t, nil, err)
	assert.True(t, ok)

	err = b.Delete(1)
	assert.Equal(t, nil, err)

	ok, err = b.Contains(1)
	assert.Equal(t, nil, err)
	assert.False(t, ok)
}
