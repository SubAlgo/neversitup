package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOddNumber(t *testing.T) {
	type testCase struct {
		caseName string
		input    []int
		expected int
	}

	testCases := []testCase{
		{caseName: "TestFindOddNumber_1", input: []int{7}, expected: 7},
		{caseName: "TestFindOddNumber_2", input: []int{0}, expected: 0},
		{caseName: "TestFindOddNumber_3", input: []int{1, 1, 2}, expected: 2},
		{caseName: "TestFindOddNumber_4", input: []int{0, 1, 0, 1, 0}, expected: 0},
		{caseName: "TestFindOddNumber_5", input: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, expected: 4},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			res := FindOddNumber(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
