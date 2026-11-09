package main

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaitGroup_WaitsForAll(t *testing.T) {
	const n = 100
	wg := New(n)

	var counter int64
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	assert.Equal(t, int64(n), atomic.LoadInt64(&counter))
}
