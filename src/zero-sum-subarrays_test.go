package main

import "testing"

func TestCountZeroSumSubarraysForCase(t *testing.T) {
	testCase := []int{0, 0, 5, 5, 0, 0}
	expectedOutput := 6
	output := countZeroSumSubarraysForCase(testCase)

	if expectedOutput != output {
		t.Errorf("Expected output to be %v, but got %v", expectedOutput, output)
	}
}
