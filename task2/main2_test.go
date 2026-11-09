package main

import (
	"reflect"
	"testing"
)

func Test_addElements(t *testing.T) {
	tests := []struct {
		name          string
		originalSlice []int
		n             int
		want          []int
	}{
		{"add 4", []int{1, 2, 3}, 4, []int{1, 2, 3, 4}},
		{"empty slice", []int{}, 4, []int{4}},
		{"nil slice", nil, 7, []int{7}},
		{"add negative", []int{1, 2}, -5, []int{1, 2, -5}},
		{"add zero", []int{9}, 0, []int{9, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addElements(tt.originalSlice, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copySlice(t *testing.T) {
	tests := []struct {
		name          string
		originalSlice []int
	}{
		{"empty slice", []int{}},
		{"nil slice", nil},
		{"single element slice", []int{42}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copySlice(tt.originalSlice); !reflect.DeepEqual(got, tt.originalSlice) {
				t.Errorf("copySlice() = %v, want %v", got, tt.originalSlice)
			}
		})
	}
}

func TestCopySlice_NotAffectOriginal(t *testing.T) {
	in := []int{1, 2, 3}
	out := copySlice(in)

	if len(out) > 0 && len(in) > 0 {
		out[0] = 100
		if in[0] == 100 {
			t.Errorf("change affects original slice")
		}
	}
}

func Test_sliceExample(t *testing.T) {
	tests := []struct {
		name          string
		originalSlice []int
		want          []int
	}{
		{"even only", []int{3, 5, 7, 33, 55}, []int{}},
		{"even + odd", []int{12, 3, 694, 32, 5, 7, 140, 33, 55}, []int{12, 694, 32, 140}},
		{"empty slice", []int{}, []int{}},
		{"nil", nil, []int{}},
		{"null + negative", []int{-45, 0, -123, -524, 90, -3, -28}, []int{0, -524, 90, -28}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceExample(tt.originalSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeElement1(t *testing.T) {
	tests := []struct {
		name          string
		originalSlice []int
		n             int
		want          []int
		wantErr       bool
	}{
		{"remove from middle", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}, false},
		{"remove first", []int{1, 2, 3}, 0, []int{2, 3}, false},
		{"remove last", []int{1, 2, 3}, 2, []int{1, 2}, false},
		{"single element slice", []int{42}, 0, []int{}, false},
		{"n out of range", []int{1, 2, 3}, 5, nil, true},
		{"negative n", []int{1, 2, 3}, -1, nil, true},
		{"empty slice", []int{}, 0, nil, true},
		{"nilslice ", nil, 0, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeElement(tt.originalSlice, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElement() got = %v, want %v", got, tt.want)
			}
		})
	}
}
