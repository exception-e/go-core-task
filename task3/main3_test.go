package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringIntMap_AddToNilMap(t *testing.T) {
	m := StringIntMap{}
	require.NotPanics(t, func() { m.Add("first", 1) })
	v, ok := m.Get("first")
	assert.True(t, ok)
	assert.Equal(t, 1, v)
}

func TestStringIntMap_Copy(t *testing.T) {
	m := New()
	m.Add("first", 1)
	m.Add("second", 2)

	copyMap := m.Copy()

	copyMap["first"] = 111
	copyMap["second"] = 222
	copyMap["third"] = 333

	v, ok := m.Get("first")
	require.True(t, ok)
	assert.Equal(t, 1, v)
	assert.False(t, m.Exists("third"))
}
