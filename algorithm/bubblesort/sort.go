package main

import (
	"fmt"
	"math/rand"
)


// 冒泡排序
func Bubble(slice []int) {
	for i := 1; i < len(slice); i++ {
		// 提前退出冒泡排序的标志
		flag := false
		for j := 0; j < len(slice)-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func main() {
	var slice []int
	for i := 1; i <= 55; i++ {
		slice = append(slice, rand.Intn(1000))
	}
	Bubble(slice)
	fmt.Println(slice)
}
