package main

import (
	"strconv"
	"strings"
)

func formatInput(rawInput string) [][]int {
	input := strings.Split(rawInput, "\n")

	var caseCount int
	caseCount, input = shiftInt(input)

	testCases := convertCases(caseCount, input)

	return testCases
}

func convertCases(caseCount int, input []string) [][]int {
	var testCases [][]int
	for i := 0; i < caseCount; i++ {
		var numberCount int
		numberCount, input = shiftInt(input)

		var testCase []int
		testCase, input = shiftCase(numberCount, input)

		testCases = append(testCases, testCase)
	}

	return testCases
}

func shiftCase(numberCount int, slice []string) ([]int, []string) {
	shiftString, shiftedSlice := slice[0], slice[1:]
	testCase := convertCase(numberCount, shiftString)

	return testCase, shiftedSlice
}

func convertCase(numberCount int, rawCase string) []int {
	splitCase := strings.Split(rawCase, " ")
	var testCase []int
	for i := 0; i < numberCount; i++ {
		var number int
		number, splitCase = shiftInt(splitCase)

		testCase = append(testCase, number)
	}

	return testCase
}

func shiftInt(slice []string) (int, []string) {
	shiftString, shiftedSlice := slice[0], slice[1:]
	shiftInt, _ := strconv.Atoi(shiftString)

	return shiftInt, shiftedSlice
}

func formatOutput(rawOutput []int) string {
	stringedOutput := convertToStrings(rawOutput)
	output := strings.Join(stringedOutput, "\n")

	return output
}

func convertToStrings(intSlice []int) []string {
	var stringSlice []string
	for _, integer := range intSlice {
		stringSlice = append(stringSlice, strconv.Itoa(integer))
	}

	return stringSlice
}
