package main

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   bool
		want1  []int
	}{
		{
			name:   "no intersection",
			slice1: []int{1, 2},
			slice2: []int{3, 4},
			want:   false,
			want1:  []int{},
		},
		{
			name:   "duplicates in slice2",
			slice1: []int{1, 2, 3},
			slice2: []int{2, 2, 2},
			want:   true,
			want1:  []int{2},
		},
		{
			name:   "duplicates in slice1",
			slice1: []int{2, 2, 2},
			slice2: []int{2},
			want:   true,
			want1:  []int{2},
		},
		{
			name:   "empty slice1",
			slice1: nil,
			slice2: []int{1, 2},
			want:   false,
			want1:  []int{},
		},
		{
			name:   "empty slice2",
			slice1: []int{1, 2},
			slice2: nil,
			want:   false,
			want1:  []int{},
		},
		{
			name:   "both slices empty",
			slice1: nil,
			slice2: nil,
			want:   false,
			want1:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := intersection(tt.slice1, tt.slice2)
			if got != tt.want {
				t.Errorf("intersection() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("intersection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
