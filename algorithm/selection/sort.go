package main

import (
	"fmt"
)

func Selection(slice []int){

	for i := 1; i < len(slice); i++ {
		min := i-1
		for j := i; j < len(slice); j++ {

			if slice[min] > slice[j] {
				min = j
			}
		}
		slice[i-1],slice[min] = slice[min],slice[i-1]
	}
}

func main()  {
	slice := []int{2,3,1,4,32,42,88,21,6,8,9,5,14,12,18,54,7,35,36,24,13}
	Selection(slice)
	fmt.Println(slice)
}