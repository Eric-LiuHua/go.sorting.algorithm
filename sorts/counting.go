package sorts

//计数排序
func CountingSort(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}
	min := nums[0]
	max := nums[0]
	res := make([]int, 0)
	for _, v := range nums {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}

	counts := make([]int, max-min+1)
	for _, v := range nums {
		counts[v-min]++
	}
	for i := 0; i < len(counts); i++ {
		tmp := counts[i]
		for j := 0; j < tmp; j++ {
			res = append(res, i+min)
		}
	}
	return res
}
