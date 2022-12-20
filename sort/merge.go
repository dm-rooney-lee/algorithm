package sort

import "golang.org/x/exp/constraints"

func Merge[T constraints.Ordered](arr []T) []T {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	left, right := Merge(arr[:mid]), Merge(arr[mid:])

	return merge(left, right)
}

func merge[T constraints.Ordered](left, right []T) []T {
	arr := make([]T, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) || j < len(right) {
		if i == len(left) {
			arr = append(arr, right[j:]...)
			break
		} else if j == len(right) {
			arr = append(arr, left[i:]...)
			break
		}

		if left[i] <= right[j] {
			arr = append(arr, left[i])
			i++
		} else {
			arr = append(arr, right[j])
			j++
		}
	}

	return arr
}
