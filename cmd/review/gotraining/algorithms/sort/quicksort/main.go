//O(nlogn)
//O(n2)
func Quicksort(nums []int, p int, r int) {
	if p < r {
		q := Partition(nums, p, r)
		Quicksort(nums, p, q-1)
		Quicksort(nums, q+1, r)
	}
}

func Partition(nums []int, p int, r int) int {
	x := nums[r]
	i := p - 1
	for j := p; j < r-1; j++ {
		if nums[j] <= x {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}