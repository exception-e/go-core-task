package main

import (
	"bytes"
	"crypto/sha256"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_hashRunesInput(t *testing.T) {
	t.Parallel()

}

func Test_makeStringOfVars(t *testing.T) {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	vars := []any{numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum}
	formatMask := []string{"%d", "%#o", "%#x", "%f", "%s", "%t", "%g"}
	t.Parallel()

	want := "420520x2a3.140000Golangtrue(1+2i)"
	got := makeStringOfVars(vars, formatMask)

	require.NotEmpty(t, got)
	assert.Equal(t, want, got)
}

func Test_printVars(t *testing.T) {

	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	vars := []any{numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum}
	formatMask := []string{"%d", "%#o", "%#x", "%f", "%s", "%t", "%g"}
	t.Parallel()
	want := []string{
		"type of 42: int",
		"type of 052: int",
		"type of 0x2a: int",
		"type of 3.140000: float64",
		"type of Golang: string",
		"type of true: bool",
		"type of (1+2i): complex64",
	}
	var buf bytes.Buffer
	printVars(&buf, vars, formatMask)
	got := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")

	require.NotEmpty(t, got)
	require.Equal(t, len(want), len(got))
	assert.Equal(t, want, got)
}

func Test_hashRunes(t *testing.T) {
	type args struct {
		runes []rune
		salt  []rune
	}
	tests := []struct {
		name string
		args args
		want [32]byte
	}{
		{name: "even", args: args{
			runes: []rune("0x2a3."),
			salt:  []rune("go-2024"),
		}, want: sha256.Sum256([]byte("0x2go-2024a3."))},

		{name: "odd", args: args{
			runes: []rune("0x2a3"),
			salt:  []rune("go-2024"),
		}, want: sha256.Sum256([]byte("0xgo-20242a3"))},
		{name: "emptyInput", args: args{
			runes: []rune(""),
			salt:  []rune("go-2024"),
		}, want: sha256.Sum256([]byte("go-2024"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, hashRunes(tt.args.runes, tt.args.salt), "hashRunes(%v, %v)", tt.args.runes, tt.args.salt)
		})
	}
}
