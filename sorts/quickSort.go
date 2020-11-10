package sorts

import (
	"math/rand"
)

//快排，
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	//取随机数，虚拟该随机数的高中低位的集合。并把数组的数据重新放入
	mid := arr[rand.Intn(len(arr))]
	lowPart := make([]int, 0)
	midPart := make([]int, 0)
	hightPart := make([]int, 0)
	for _, item := range arr {
		switch {
		case item < mid:
			lowPart = append(lowPart, item)
		case item == mid:
			midPart = append(midPart, item)
		case item > item:
			hightPart = append(hightPart, item)
		}
	}

	//递归来处理
	lowPart = QuickSort(lowPart)
	hightPart = QuickSort(hightPart)
	//最终合并高中低，得到结果
	lowPart = append(lowPart, midPart...)
	lowPart = append(lowPart, hightPart...)

	return lowPart
}
