package heap

import "fmt"

// 快速排序
func Quick(slice []int, left int, right int) {
	if left >= right {
		return
	}
	p := Partition(slice, left, right)
	Quick(slice, left, p-1)
	Quick(slice, p+1, right)
}


func Partition(slice []int, left int, right int) (p int){
	pivot := slice[right]
	i := left
	for j := left; j < right; j++ {
		if slice[j] < pivot {
			if i != j {
				slice[i],slice[j] = slice[j],slice[i]
			}
			i++
		}
	}
	slice[right],slice[i] = slice[i],slice[right]
	return i
}


func main() {
	slice := []int{2, 3, 1, 4, 32, 42, 88, 21, 6, 8, 9, 5, 14, 12, 18, 54, 7, 35, 36, 24, 13}
	Quick(slice, 0, len(slice)-1)
	fmt.Println(slice)
}
