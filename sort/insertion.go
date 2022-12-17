package sort

func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp, j := arr[i], i-1
		for ; j >= 0 && temp < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
}
