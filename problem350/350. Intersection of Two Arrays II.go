package problem350

import (
	"fmt"
	"sort"
)

func Judge() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	//nums1 := []int{4, 9, 5}       // 4 5 9
	//nums2 := []int{9, 4, 9, 8, 4} // 4 4 8 9 9
	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	index1 := 0
	index2 := 0

	result := make([]int, 0)

	for {
		if index1 == len(nums1)-1 || index2 == len(nums2)-1 {
			if nums1[index1] == nums2[index2] {
				result = append(result, nums1[index1])
				break
			} else if index1 != len(nums1)-1 {
				index1++
			} else if index2 != len(nums2)-1 {
				index2++
			} else {
				break
			}
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else if nums1[index1] < nums2[index2] {
			index1++
		} else {
			result = append(result, nums1[index1])
			index1++
			index2++
		}
	}

	return result
}
