package sorts

//冒泡排序
func BubbleSort(nums []int) []int {
	var len int = len(nums)
	for i := 0; i < len; i++ {
		for j := len - 1; j > i; j-- {
			if nums[j] < nums[i] {
				nums[j] = nums[j] ^ nums[i]
				nums[i] = nums[j] ^ nums[i]
				nums[j] = nums[j] ^ nums[i]
			}
		}
	}
	return nums
}

//
func BubbleSortA(nums []int) []int {
	var len int = len(nums)
	for i := 0; i < len; i++ {
		for j := len - 1; j > i; j-- {
			if nums[j] < nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
	return nums
}
