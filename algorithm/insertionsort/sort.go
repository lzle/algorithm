package main

import (
	"fmt"
	"math/rand"
)

// 插入排序
func Insertion(slice []int) {

	for i := 1; i < len(slice); i++ {
		v := slice[i]
		j := i - 1
		for ; j >= 0; j-- {
			if slice[j] > v {
				slice[j+1] = slice[j]
			}else {
				break
			}
		}
		slice[j+1]= v
	}
}

func main() {
	var slice []int
	for i := 1; i <= 50; i++ {
		slice = append(slice, rand.Intn(1000))
	}
	Insertion(slice)
	fmt.Println(slice)
}

