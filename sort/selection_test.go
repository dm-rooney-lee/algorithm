package sort

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestSelection(t *testing.T) {
	type tc[T constraints.Ordered] struct {
		name string
		arr  []T
		want []T
	}

	tests := []tc[int]{
		{
			name: "success case: even len arr",
			arr:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "success case: odd len arr",
			arr:  []int{7, 5, 3, 10, 1, 2, 9, 8, 4},
			want: []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Selection(tt.arr)
			if got := tt.arr; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Selection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := setupBench(testSize)
		Selection(arr)
	}
}
