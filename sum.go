package module01

// Sum will sum up all of the numbers passed
// in and return the result
// func Sum(numbers []int) int {
// 	// Declare a sum value to 0 to be updated
// 	runningSum := 0
// 	// Iterate through the collection
// 	for _, value := range numbers {
// 		// Addition Reassign Sum
// 		runningSum += value
// 	}

// 	return runningSum
// }
// Recursive Approach
func Sum(numbers []int) int {
	// Base Case
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}
