package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pipe(t *testing.T) {
	ch := make(chan uint8)
	values := []uint8{1, 2, 3, 4, 5}
	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()

	var got []float64
	for v := range pipe(ch) {
		got = append(got, v)
	}

	want := []float64{
		1,
		8,
		27,
		64,
		125,
	}
	assert.Equal(t, want, got)
}
