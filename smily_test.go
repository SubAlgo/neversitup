package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmilyFace(t *testing.T) {
	type testCase struct {
		caseName string
		input    []string
		expected int
	}

	testCases := []testCase{
		{caseName: "TestCountSmilyFace_1", input: []string{":)", ";(", ";}", ":-D"}, expected: 2},
		{caseName: "TestCountSmilyFace_2", input: []string{";D", ":-(", ":-)", ";~)"}, expected: 3},
		{caseName: "TestCountSmilyFace_3", input: []string{";]", ":[", ";*", ":$", ";-D"}, expected: 1},
	}

	for _, tc := range testCases {
		t.Run(tc.caseName, func(t *testing.T) {
			res := CountSmilyFace(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
