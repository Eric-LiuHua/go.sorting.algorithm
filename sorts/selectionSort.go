package sorts

//选择排序，先找到最小数字，在交换
func SelectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		//找到右边最小值的坐标
		tmp := i
		for j := i + 1; j < len(arr); j++ {
			if arr[tmp] > arr[j] {
				tmp = j
			}
		}
		//开始交换，这样使用异或。那就需要判断两者不想等，否则出现0
		arr[tmp], arr[i] = arr[i], arr[tmp]

	}
	return arr
}
