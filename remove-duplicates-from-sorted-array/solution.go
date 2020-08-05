package main

// RemoveDuplicates returns the number of unique elements in the nums slice
func RemoveDuplicates(nums []int) int {
	m := make(map[int]int)
	for _, item := range nums {
		_, ok := m[item]
		if !ok {
			m[item] = 1
			nums[len(m)-1] = item
		}
	}
	return len(m)
}
