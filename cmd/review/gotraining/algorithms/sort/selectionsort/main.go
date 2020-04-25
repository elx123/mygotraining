func selects(nums []int) {
	for k, _ := range nums {
		index := k
		for j := k; j < len(nums); j++ {
			if nums[j] < nums[index] {
				index = j
			}
		}
		if index != k {
			nums[k], nums[index] = nums[index], nums[k]
		}
	}
}