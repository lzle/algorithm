package main

import "fmt"

// weight:物品重量，n:物品个数，w:背包可承载重量
func knapsack(weight []int, n int, w int) int{
	states := make([][]bool, n)  // 默认值false
	for i:=0; i<n; i++ {
		states[i] = make([]bool, w+1)
	}

	states[0][0] = true
	if weight[0] <= w {
		states[0][weight[0]] = true
	}

	for i:=1; i<n; i++ {
		// 不放
		for j := 0; j <= w; j++ {
			if states[i-1][j] {
				states[i][j] = true
			}
		}
		// 放
		for j := 0; j <= w-weight[i]; j++ {
			if states[i-1][j] {
				states[i][j+weight[i]] = true
			}
		}
	}

	for i:=w; i>=0; i-- {
		if states[n-1][i] {
			return i
		}
	}

	return 0
}

func main()  {
	wight := []int{2,2,5,6,11}
	n := 5
	w := 12

	ret := knapsack(wight, n , w)
	fmt.Println(ret)
}