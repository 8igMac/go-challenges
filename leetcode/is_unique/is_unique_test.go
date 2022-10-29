package main

import "testing"

type testCase struct {
	str    string
	expect bool
}

func TestIsUnique(t *testing.T) {
	testCases := []testCase{
		{"abc", true},
		{"aac", false},
		{".", true},
		{"a als;dfjasldfk ac", false},
	}

	for _, testCase := range testCases {
		if result := IsUnique(testCase.str); result != testCase.expect {
			t.Fatalf("Str: %v, expect: %v, got %v.", testCase.str, testCase.expect, result)
		}
	}

}
