package main

func countZeroSumSubarraysForCases(testCases [][]int) []int {
	var outputCases []int
	for _, testCase := range testCases {
		outputCase := countZeroSumSubarraysForCase(testCase)
		outputCases = append(outputCases, outputCase)
	}

	return outputCases
}

func countZeroSumSubarraysForCase(testCase []int) int {
	var count int
	var sum int
	sums := map[int][]int{0: []int{-1}}
	for i, number := range testCase {
		sum += number

		if sums[sum] != nil {
			count += len(sums[sum])
			sums[sum] = append(sums[sum], i)
		} else {
			sums[sum] = []int{i}
		}
	}

	return count
}
