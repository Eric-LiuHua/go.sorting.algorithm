package sorts

//插入排序
func InsertSort(arr []int) []int {
	for out := 0; out < len(arr); out++ {
		tmp := arr[out]
		in := out
		//先把当前位置到值记录，
		//判断大于当前值的，都往后移动，在把当前值插入移动后的坐标
		for ; in > 0 && arr[in-1] >= tmp; in-- {
			arr[in] = arr[in-1]
		}
		arr[in] = tmp

	}
	return arr
}
