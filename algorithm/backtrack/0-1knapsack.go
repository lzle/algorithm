package backtrack

var maxW int

// cw表示当前已经装进去的物品的重量和；
//i表示考察到哪个物品了；
//w背包重量；items表示每个物品的重量；
//n表示物品个数
//假设背包可承受重量100，物品个数10，物品重量存储在数组a中，那可以这样调用函数：
//f(0, 0, a, 10, 100)
func knapsack(i int, cw int, items []int, n int, w int) {
	if cw == w || i == n {
		if cw > maxW {
			maxW = cw
		}
		return
	}
	// 不放
	knapsack(i+1, cw, items, n, w)
	// 放
	if cw+items[i] <= w {
		knapsack(i+1, cw+items[i], items, n, w)
	}
}
