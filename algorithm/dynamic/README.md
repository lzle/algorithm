# 动态规划

动态规划比较适合用来求解最优问题，比如求最大值、最小值等等。它可以非常显著地降低时间复杂度，提高代码的执行效率。

### 0-1 背包问题

关于这个问题，使用回溯的解决方法，枚举所有情况，然后找出满足条件的最大值。不过，回溯算法的复杂度比较高，是指数级别的。

``` go
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
```

通过上面例子可以找到规律。

<img src="https://static001.geekbang.org/resource/image/42/ea/42ca6cec4ad034fc3e5c0605fbacecea.jpg" width=500>

发现图中 f(2, 2) 和 f(3,4) 都被重复计算了两次，可以利用备忘录进行避免冗余计算。

```` go
func knapsack(i int, cw int, items []int, n int, w int) {
	if cw == w || i == n {
		if cw > maxW {
			maxW = cw
		}
		return
	}
	if mem[i][cw] {
		return
	}
	mem[i][cw] = true
	// 不放
	knapsack(i+1, cw, items, n, w)
	// 放
	if cw+items[i] <= w {
		knapsack(i+1, cw+items[i], items, n, w)
	}
}
````

这种解决方法非常好。实际上，它已经跟动态规划的执行效率基本上没有差别。

动态规划把求解过程分为 n 个阶段，每个阶段会决策一个物品是否放到背包中。每一层重复的状态（节点）合并，
只记录不同的状态，然后基于上一层的状态集合，来推导下一层的状态集合。

通过合并每一层重复的状态，这样就保证每一层不同状态的个数都不会超过 w 个（w 表示背包的承载重量),避免指数增长。

用一个二维数组 states[n][w+1]，来记录每层可以达到的不同状态。第 0 个（下标从 0 开始编号）物品的重量是 2，
要么装入背包，要么不装入背包，决策完之后，会对应背包的两种状态，背包中物品的总重量是 0 或者 2。
我们用 states[0][0]=true 和 states[0][2]=true 来表示这两种状态。

<img src="https://static001.geekbang.org/resource/image/bb/7e/bbbb934247219db8299bd46dba9dd47e.jpg" width=500>

把所有物品考察完，整个 states 状态数组就都计算好了。

``` go
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
```

用回溯算法解决这个问题的时间复杂度 O(2^n)，是指数级的。用动态规划的时间复杂度是 O(n*w)。

还有更节省空间的方式

```go
// weight:物品重量，n:物品个数，w:背包可承载重量
func knapsack(weight []int, n int, w int) int{
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
```
j 需要从大到小来处理。








