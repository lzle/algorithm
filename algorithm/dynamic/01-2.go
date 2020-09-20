package main

import "fmt"

// weight:物品重量，n:物品个数，w:背包可承载重量
func knapsack2(weight []int, n int, w int) int{
	states := make([]bool, w+1)  // 默认值false
	states[0] = true
	if weight[0] < w {
		states[weight[0]] = true
	}
	for i:=1; i<n; i++ {
		for j:= w-weight[i]; j>=0; j-- {
			if states[j]{
				states[j+weight[i]] = true
			}
		}
	}
	for i:=w; i>=0; i-- {
		if states[i] {
			return i
		}
	}
	return 0
}

func main()  {
	wight := []int{2,2,5,6,11}
	n := 5
	w := 12
	ret := knapsack2(wight, n , w)
	fmt.Println(ret)
}
