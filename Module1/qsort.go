package main

import "fmt" //это пакет

var arr []int

func swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func less(i int, j int) bool {
	if arr[i] >= arr[j] {
		return false
	}
	return true
}
func Partition(swap func(i, j int), less func(i, j int) bool, low int, high int) int {
	i := low
	j := high
	for j < high {
		if less(arr[j], arr[high]) {
			swap(i, j)
			i++
		}
		j++
	}
	swap(i, high)
	return i
}
func QuickSortRec(less func(i, j int) bool, swap func(i, j int), low int, high int) {
	if less(low, high) {
		q := Partition(swap, less, low, high)
		QuickSortRec(less, swap, low, q-1)
		QuickSortRec(less, swap, q+1, high)
	}
}
func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	QuickSortRec(less, swap, 0, n-1)
}

func main() {
	//fmt.Println("This is IT!")
	var n int
	fmt.Scan(&n)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	qsort(n, less, swap)
	fmt.Println(arr)
}

