package main

import "testing"

func TestAdd(t *testing.T) {
	actualValue := add(2, 5)
	expectedValue := 7

	if actualValue != expectedValue {
		t.Errorf("got %v, wanted %v", actualValue, expectedValue)
	}
}
