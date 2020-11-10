package sorts

//对比数组
func CompareSlices(x []int, y []int) (int, bool) {

	if len(x) != len(y) {
		return -1, false
	}

	for pos := range x {
		if x[pos] != y[pos] {
			return pos, false
		}
	}
	return 0, true
}
