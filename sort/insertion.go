package sort

import "golang.org/x/exp/constraints"

func Insertion[T constraints.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		temp, j := arr[i], i-1
		for ; j >= 0 && temp < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
}
