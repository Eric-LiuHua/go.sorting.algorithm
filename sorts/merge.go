package sorts

//归并排序
func Merge(arr []int) []int {
	arr_len := len(arr)
	if arr_len > 1 {
		var tmp []int = make([]int, arr_len)
		MergePart(arr, tmp, 0, arr_len-1)
	}
	return arr
}

//arr 原数组
//tmp 辅助数组
//left 左边
//right 右边
func MergePart(arr []int, tmp []int, left int, right int) {
	if left < right {
		//取中位
		mid := left + ((right - left) >> 1)
		//i 左边最大位，i 右边最小位
		var i int = left
		var j int = mid + 1
		var cur int = 0

		MergePart(arr, tmp, left, mid)
		MergePart(arr, tmp, j, right)

		for {
			if i <= mid && j <= right {
				//按大小，按顺序放入辅助数组。
				if arr[i] < arr[j] {
					tmp[cur] = arr[i]
					i += 1
				} else {
					tmp[cur] = arr[j]
					j += 1
				}
				cur += 1
			} else if i <= mid {
				//把左边还未取完的数据，继续取出来
				tmp[cur] = arr[i]
				cur += 1
				i += 1
			} else if j <= right {
				//把右边还未取完的数据，继续取出来
				tmp[cur] = arr[j]
				cur += 1
				j += 1
			} else {
				break
			}
		}

		//辅助数组排序结束
		//复原到原数组中，范围就是left -right
		for ti := 0; ti < cur; ti++ {
			arr[left+ti] = tmp[ti]
		}
	}
}
