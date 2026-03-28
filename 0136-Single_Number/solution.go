package main

func SingleNumber(nums []int) int {
	result := 0
	freqMap := make(map[int]int8, len(nums)/2)
	for _, value := range nums {
		freqMap[value]++
		if freqMap[value] == 1 {
			result += value
		} else {
			result -= value
		}
	}
	return result
}

func SingleNumberWithXOR(nums []int) int {
	result := 0
	for _, value := range nums {
		result ^= value
	}
	return result
}

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
