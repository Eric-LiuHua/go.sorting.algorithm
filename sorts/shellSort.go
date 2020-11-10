package sorts

//
func ShellSort(arr []int) []int {
	for i := (len(arr) >> 1); i > 0; i = i >> 1 {
		for j := i; j < len(arr); j++ {
			for t := i; t >= i && arr[t-i] > arr[t]; t -= i {
				arr[t], arr[t-i] = arr[t-i], arr[t]
			}

		}
	}
	return arr
}

func ShellSortA(arr []int) []int {
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
	return arr
}
