package main

import "testing"

func TestFooer(t *testing.T) {
	expect := false
	result := false
	if result != true {
		t.Errorf("Result was incorrect, got: %v, want: %v", result, expect)
	}
}
