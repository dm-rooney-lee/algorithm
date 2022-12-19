package sort

func Heap(arr []int) {
	last := len(arr) - 1
	for i := (last - 1) / 2; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	for size := last; size > 0; size-- {
		arr[0], arr[size] = arr[size], arr[0]
		heapify(arr, 0, size)
	}
}

func heapify(arr []int, root, size int) {
	max := root
	l, r := root*2+1, root*2+2
	if l < size && arr[l] > arr[max] {
		max = l
	}
	if r < size && arr[r] > arr[max] {
		max = r
	}

	if max != root {
		arr[root], arr[max] = arr[max], arr[root]
		heapify(arr, max, size)
	}
}
