func bub(nums []int) {
	for k, _ := range nums {
		for j := k; j < len(nums); j++ {
			if nums[j] > nums[k] {
				nums[j], nums[k] = nums[k], nums[j]
			}
		}
	}
}