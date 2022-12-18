package sort

func Merge(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	left, right := Merge(arr[:mid]), Merge(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	arr := make([]int, 0, len(left)+len(right))
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
