package main

import (
	"fmt"
)

// 归并排序
func Merge(slice []int, left int, right int) {

	if left < right {
		mid := (left + right) / 2
		Merge(slice, left, mid)
		Merge(slice, mid+1, right)

		var temp []int
		i := left
		j := mid + 1
		for i <= mid && j <= right {
			if slice[i] <= slice[j] {
				temp = append(temp, slice[i])
				i++
				continue
			}
			temp = append(temp, slice[j])
			j++
		}
		if i <= mid {
			temp = append(temp, slice[i:mid+1]...)
		}
		if j <= right {
			temp = append(temp, slice[j:right+1]...)
		}

		for i, v := range temp {
			slice[left+i] = v
		}
	}
}

func main() {
	slice := []int{2, 3, 1, 4, 32, 42, 88, 21, 6, 8, 9, 5, 14, 12, 18, 54, 7, 35, 36, 24, 13}
	Merge(slice, 0, len(slice)-1)
	fmt.Println(slice)
}
