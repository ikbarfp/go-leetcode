package main_test

import (
	implPkg "github.com/ikbarfp/go-leetcode/0136-Single_Number"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	exp  int
	data []int
}

var testCases = []testCase{
	{1, []int{2, 2, 1}},
	{4, []int{4, 1, 2, 1, 2}},
	{1, []int{1}},
	{1, []int{4, 1, 4}},
}

func TestSingleNumber_ShouldBeCorrect(t *testing.T) {
	for _, tc := range testCases {
		res := implPkg.SingleNumber(tc.data)
		assert.Equal(t, tc.exp, res)
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		implPkg.SingleNumber(testCases[0].data)
	}
}

func TestSingleNumberWithXOR_ShouldBeCorrect(t *testing.T) {
	for _, tc := range testCases {
		res := implPkg.SingleNumberWithXOR(tc.data)
		assert.Equal(t, tc.exp, res)
	}
}

func BenchmarkSingleNumberWithXOR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		implPkg.SingleNumberWithXOR(testCases[0].data)
	}
}
