package sort

func Heap(arr []int) {
	last := len(arr) - 1
	for root := (last - 1) / 2; root >= 0; root-- {
		heapify(arr, root, len(arr))
	}

	for end := last; end > 0; end-- {
		arr[0], arr[end] = arr[end], arr[0]
		heapify(arr, 0, end)
	}
}

func heapify(arr []int, root, end int) {
	max := root
	l, r := root*2+1, root*2+2
	if l < end && arr[l] > arr[max] {
		max = l
	}
	if r < end && arr[r] > arr[max] {
		max = r
	}

	if max != root {
		arr[root], arr[max] = arr[max], arr[root]
		heapify(arr, max, end)
	}
}
