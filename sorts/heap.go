package sorts

//从上面对于堆排序的叙述我们知道，进行一次堆排序，我们要解决两个问题：
//1、如何初始化一个堆
//2、如何在输出堆顶元素之后，调整堆内元素，使其再次形成一个堆。

//最大堆结构
type MaxHeap struct {
	HeapSlice []int
	HeapSize  int
}

func (h MaxHeap) MaxHeapify(i int) {
	l, r := (i<<1)+1, (i<<1)+2
	max := i
	if l < h.HeapSize && h.HeapSlice[l] > h.HeapSlice[max] {
		max = l
	}
	if r < h.HeapSize && h.HeapSlice[r] > h.HeapSlice[max] {
		max = r
	}
	if max != i {
		h.HeapSlice[i], h.HeapSlice[max] = h.HeapSlice[max], h.HeapSlice[i]
		h.MaxHeapify(max)
	}

}

func buildMaxHeap(slice []int) MaxHeap {
	h := MaxHeap{
		HeapSlice: slice,
		HeapSize:  len(slice),
	}

	for i := len(slice) >> 1; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h
}

//堆排序
func HeapSort(slice []int) []int {
	h := buildMaxHeap(slice)
	for i := len(h.HeapSlice) - 1; i >= 1; i-- {
		h.HeapSlice[i], h.HeapSlice[0] = h.HeapSlice[0], h.HeapSlice[i]
		h.HeapSize--
		h.MaxHeapify(0)

	}
	return h.HeapSlice
}
