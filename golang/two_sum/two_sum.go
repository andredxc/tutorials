/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order. (I changed this to keep testing simpler)

*/

package two_sum

func twoSum(numbers []int, target int) []int {

	var remainder int

	numberMap := make(map[int]int, len(numbers))

	// Add all numbers to the number map for easy access
	for i, num := range numbers {
		if _, ok := numberMap[num]; !ok {
			numberMap[num] = i
		}
	}

	// Find the remainder in the map
	for i, num := range numbers {
		remainder = target - num
		if j, ok := numberMap[remainder]; ok {
			if i != j {
				// Can't use the same element twice
				if i < j {
					return []int{i, j}
				} else {
					return []int{j, i}
				}
			}
		}
	}
	return nil
}
