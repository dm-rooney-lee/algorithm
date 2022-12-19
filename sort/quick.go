package sort

func Quick(arr []int) {
	quick(arr, 0, len(arr)-1)
}

func quick(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quick(arr, low, pi-1)
		quick(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot, i := arr[low], low
	for j := low + 1; j <= high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[low], arr[i] = arr[i], arr[low]

	return i
}
