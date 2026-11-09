package main

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_mergeChanns(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	go func() {
		defer close(ch3)
		ch3 <- 5
		ch3 <- 6
	}()

	var got []int
	deadline := time.After(2 * time.Second)

	for v := range mergeChanns(ch1, ch2, ch3) {
		got = append(got, v)
		select {
		case <-deadline:
			t.Fatal("Output failed")
		default:
		}
	}

	sort.Ints(got)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, got)
}
