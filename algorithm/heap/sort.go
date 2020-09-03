package main

import "fmt"

// 堆排序
func heap(slice []int, n int) {
	buildHeap(slice, n)
	for k := n; k > 1; k-- {
		slice[1], slice[k] = slice[k], slice[1]
		heapify(slice, k-1, 1)
	}
}

// 建堆
func buildHeap(slice []int, n int) {
	for i := n / 2; i >= 1; i-- {
		heapify(slice, n, i)
	}
}

// 堆化
func heapify(slice []int, n int, i int) {
	for {
		max := i
		if i*2 <= n && slice[max] < slice[i*2] {
			max = i * 2
		}
		if i*2+1 <= n && slice[max] < slice[i*2+1] {
			max = i*2 + 1
		}
		if max == i {
			break
		}
		slice[max], slice[i] = slice[i], slice[max]
		i = max
	}
}

func main() {
	slice := []int{0, 2, 3, 1, 4, 32, 42, 88, 21, 6, 8, 9, 5, 14, 12, 18, 54, 7, 35, 36, 24, 13}
	heap(slice, len(slice)-1)
	fmt.Println(slice)
}
