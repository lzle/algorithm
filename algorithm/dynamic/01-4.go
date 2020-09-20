package  main

import "fmt"

// weight:物品重量，value:物品价值，n:物品个数，w:背包可承载重量
func knapsack4(weight []int,value []int, n int, w int) int{
	states := make([]int, w+1)  // 默认值false
	for i:=0; i<w+1; i++ {
		states[i] = -1
	}
	states[0] = 0
	if weight[0] < w {
		states[weight[0]] = value[0]
	}

	for i:=1; i<n; i++ {
		// 不放
		for j := w - weight[i]; j >= 0; j-- {
			if states[j] >= 0 {
				v := states[j] + value[i]
				if v > states[j+weight[i]] {
					states[j+weight[i]] = v
				}
			}

		}
	}
	fmt.Println(states)
	maxV := -1
	for i:=w; i>=0; i-- {
		if states[i] > maxV{
			maxV = states[i]
		}
	}

	return maxV
}

func main()  {
	wight := []int{1,2,5,6,6,11}
	value := []int{1,2,5,6,7,11}
	n := 6
	w := 12
	ret := knapsack4(wight,value, n , w)
	fmt.Println(ret)
}
