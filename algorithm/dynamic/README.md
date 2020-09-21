# 动态规划

动态规划把问题分解为n个阶段，每个阶段决策一个事件，只记录不同的状态，然后基于上一层的状态集合，
来推导下一层的状态集合。通过把相同状态进行合并，来降低时间复杂度。

动态规划比较适合用来求解最优问题，比如求最大值、最小值等等。它可以非常显著地降低时间复杂度，提高代码的执行效率。

### 0-1 背包问题

#### 回溯

枚举所有情况，然后找出满足条件的最大值。回溯算法的复杂度为 O(2^n)，指数级别。

``` go
//cw表示当前已经装进去的物品的重量和；i表示考察到哪个物品了；w背包重量；items表示每个物品的重量；n表示物品个数
//假设背包可承受重量100，物品个数10，物品重量存储在数组a中，那可以这样调用函数：f(0, 0, a, 10, 100)
func knapsack(i int, cw int, items []int, n int, w int) {
    if cw == w || i == n {
        if cw > maxW {
            maxW = cw
        }
        return
    }
    // 选择不装第i个物品
    knapsack(i+1, cw, items, n, w)
    // 选择装第i个物品
    if cw+items[i] <= w {
        knapsack(i+1, cw+items[i], items, n, w)
    }
}
```

解决问题的过程，转化为下图。

<img src="https://static001.geekbang.org/resource/image/42/ea/42ca6cec4ad034fc3e5c0605fbacecea.jpg" width=400>

可以发现规律，图中 f(2, 2) 和 f(3,4) 都被重复计算了两次。可以利用备忘录的方式，避免第二次重复计算。

```` Go
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
    // 选择不装第i个物品
    knapsack(i+1, cw, items, n, w)
    // 选择装第i个物品
    if cw+items[i] <= w {
	    knapsack(i+1, cw+items[i], items, n, w)
    }
}
````

这种解决方法非常好。实际上，它已经跟动态规划的执行效率基本上没有差别。

#### 动态规划

把求解过程分为 n 个阶段，每个阶段会决策一个物品是否放到背包中，记录状态。把一层重复的状态合并，
这样就保证每一层不同状态的个数都不会超过 w 个（w 表示背包的承载重量)，避免指数增长。时间复杂度是 O(n*w)。

用一个二维数组 states[n][w+1]，来记录每层可以达到的不同状态。

例如，第 0 个 物品的重量是 2，要么装入背包，要么不装入背包，决策完之后，会对应背包的两种状态，
背包中物品的总重量是 0 或者 2。用 states[0][0]=true 和 states[0][2]=true 来表示。

把所有物品考察完，整个 states 状态数组就都计算好了。

<img src="https://static001.geekbang.org/resource/image/bb/7e/bbbb934247219db8299bd46dba9dd47e.jpg" width=600>

处理过程。

``` Go
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

程序还可以优化，把states变成一维数组，降低空间复杂度。[01-2.go](https://github.com/lzle/algorithm/blob/master/algorithm/dynamic/01-2.go)


### 0-1 背包问题升级

我们刚刚讲的背包问题，只涉及背包重量和物品重量。我们现在引入物品价值这一变量。对于一组不同重量、不同价值、不可分割的物品，我们选择将某些物品装入背包，
在满足背包最大重量限制的前提下，背包中可装入物品的总价值最大是多少呢？

#### 回溯

``` Go
// 添加cv、maxV变量
func knapsack(i int, cw int, cv int, items []int, n int, w int) {
    if cw == w || i == n {
        if cv > maxV {
            maxV = cv
        }
        return
    }
    // 选择不装第i个物品
    knapsack(i+1, cw, items, n, w)
    // 选择装第i个物品
    if cw+items[i] <= w {
        knapsack(i+1, cw+items[i], cv+value[i], items, n, w)
    }
}
```

针对上面的代码，我们还是照例画出递归树。

<img src="https://static001.geekbang.org/resource/image/bf/3f/bf0aa18f367db1b8dfd392906cb5693f.jpg" width=400>

此时这个问题就没法再用"备忘录"解决，代码无法继续优化。

#### 动态规划

```` Go
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
    maxV := -1
    for i:=w; i>=0; i-- {
        if states[i] > maxV{
            maxV = states[i]
        }
    }
    return maxV
}
````
使用动态规划，时间复杂度降低为 O(n*w)，利用一维数组，降低空间复杂度。



