package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManipulate(t *testing.T) {

	type testModule struct {
		caseName string
		input    []string
		expected []string
	}

	testModules := []testModule{
		{caseName: "TestManipulate_1", input: []string{"a"}, expected: []string{"a"}},
		{caseName: "TestManipulate_2", input: []string{"a", "b"}, expected: []string{"ab", "ba"}},
		{caseName: "TestManipulate_3", input: []string{"a", "b", "c"}, expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{caseName: "TestManipulate_4", input: []string{"a", "a", "b", "b"}, expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}

	for _, tm := range testModules {
		t.Run(tm.caseName, func(t *testing.T) {
			res := Manipulate(tm.input)
			assert.Equal(t, tm.expected, res)
		})
	}
}
