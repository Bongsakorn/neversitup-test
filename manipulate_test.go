package main

import (
	"testing"
)

func TestManipulateCase(t *testing.T) {
	testCase := map[string]int{
		"a":      1,
		"ab":     2,  // 2! = 2
		"abc":    6,  // 3! = 6
		"aabb":   6,  // 4! / 2! = 6
		"aabbcc": 90, // 6! / 2! / 2! = 90
	}

	for k, v := range testCase {
		results := Manipulate(k)

		if v != len(results) {
			t.Errorf("Expect %d but got %d", v, len(results))
		}
	}
}
