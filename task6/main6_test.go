package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate_RandGenOutputNumbers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := randGen(ctx)

	seen := make(map[int]struct{})
	for i := 0; i < 100; i++ {
		select {
		case v, ok := <-ch:
			require.True(t, ok)
			seen[v] = struct{}{}
		case <-time.After(time.Second):
			t.Fatalf("Chan is empty")
		}
	}
	assert.Greater(t, len(seen), 1, "randGen invalid - all numbers equal")
}
