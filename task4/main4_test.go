package main

import (
	"reflect"
	"testing"
)

func Test_sliceDiff(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []string
		slice2 []string
		want   []string
	}{
		{
			name:   "empty slice2",
			slice1: []string{"apple", "banana"},
			slice2: nil,
			want:   []string{"apple", "banana"},
		},
		{
			name:   "empty slice1",
			slice1: nil,
			slice2: []string{"apple"},
			want:   []string{},
		},
		{
			name:   "both slices empty",
			slice1: nil,
			slice2: nil,
			want:   []string{},
		},
		{
			name:   "slice2 includes entire slice1",
			slice1: []string{"apple", "banana"},
			slice2: []string{"apple", "banana", "cherry"},
			want:   []string{},
		},
		{
			name:   "empty string in slice is valid",
			slice1: []string{"", "apple"},
			slice2: []string{"apple"},
			want:   []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceDiff(tt.slice1, tt.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
