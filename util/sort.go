package util

type Sortable interface {
	Len() int
	Swap(int, int)
	GetKey(int) int
	Get(int) interface{}
	Set(int, interface{})
}

func QuickSort(data Sortable) {
	quickSort(data, 0, data.Len()-1)
}
func BubbleSort(array Sortable) {
	for i := 0; i<array.Len(); i++ {
		for j := 0; j<array.Len()-i-1; j++ {
			if array.GetKey(j) > array.GetKey(j+1) {
				array.Swap(j, j+1)
			}
		}
	}
}
func quickSort(array Sortable, left int, right int) {
	if left < right {
		data := array.Get(left)
		key := array.GetKey(left)
		low := left
		high := right
		for low < high {
			for low < high && array.GetKey(high) > key {
				high--
			}
			array.Set(low, array.Get(high))
			for low < high && array.GetKey(low) < key {
				low++
			}
			array.Set(high, array.Get(low))
		}
		array.Set(low, data)
		quickSort(array, left, low-1);
		quickSort(array, low+1, right);
	}
}