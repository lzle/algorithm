package  main

import "fmt"

// weight:物品重量，value:物品价值，n:物品个数，w:背包可承载重量
func knapsack3(weight []int,value []int, n int, w int) int{
	states := make([][]int, n)  // 默认值false
	for i:=0; i<n; i++ {
		states[i] = make([]int, w+1)
	}
	for i:=0; i<n; i++ {
		for j:=0; j<w+1; j++ {
			states[i][j] = -1
		}
	}

	states[0][0] = 0
	if weight[0] <= w {
		states[0][weight[0]] = value[0]
	}

	for i:=1; i<n; i++ {
		// 不放
		for j := 0; j <= w; j++ {
			if states[i-1][j] >= 0  {
				states[i][j] = states[i-1][j]
			}
		}
		// 放
		for j := 0; j <= w-weight[i]; j++ {
			if states[i-1][j] >= 0  {
				v := states[i-1][j] + value[i]
				if (v > states[i][j+weight[i]]) {
					states[i][j+weight[i]] = v
				}
			}
		}
	}
	fmt.Println(states[n-1])
	maxV := -1
	for i:=w; i>=0; i-- {
		if states[n-1][i] > maxV{
			maxV = states[n-1][i]
		}
	}

	return maxV
}

func main()  {
	wight := []int{1,2,5,6,6,11}
	value := []int{1,2,5,6,7,11}
	n := 6
	w := 12

	ret := knapsack3(wight,value, n , w)
	fmt.Println(ret)
}
