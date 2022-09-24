package median_of_sorted_arrays

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
https://leetcode.com/problems/median-of-two-sorted-arrays/
*/

// UNFINISHED

func medianOfSortedArrays(a1, a2 []int) float32 {

	var middleIndex int

	numElements := len(a1) + len(a2)
	middleIndex = int(numElements / 2)

	if numElements%2 == 0 {
		middleIndex -= 1
	}

	indexArray1 := 0
	indexArray2 := 0
	for i := 0; i < numElements; i++ {
		value1 := 0
		if indexArray1 < len(a1) {
			value1 = a1[indexArray1]
		}

		value2 := 0
		if indexArray2 < len(a2) {
			value2 = a2[indexArray2]
		}

	}

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
