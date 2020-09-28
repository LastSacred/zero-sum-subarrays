package main

import "fmt"

func main() {
	testCases := formatInput(input)
	rawOutput := countZeroSumSubarraysForCases(testCases)
	output := formatOutput(rawOutput)

	fmt.Println(output)
}
