# 动态规划

动态规划比较适合用来求解最优问题，比如求最大值、最小值等等。它可以非常显著地降低时间复杂度，提高代码的执行效率。

### 0-1 背包问题

#### 回溯的方式

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

#### 动态规划的方式

动态规划把求解过程分为 n 个阶段，每个阶段会决策一个物品是否放到背包中。每一层重复的状态（节点）合并，
只记录不同的状态，然后基于上一层的状态集合，来推导下一层的状态集合。

通过合并每一层重复的状态，这样就保证每一层不同状态的个数都不会超过 w 个（w 表示背包的承载重量),避免指数增长。

用一个二维数组 states[n][w+1]，来记录每层可以达到的不同状态。第 0 个（下标从 0 开始编号）物品的重量是 2，
要么装入背包，要么不装入背包，决策完之后，会对应背包的两种状态，背包中物品的总重量是 0 或者 2。
我们用 states[0][0]=true 和 states[0][2]=true 来表示这两种状态。

<img src="https://static001.geekbang.org/resource/image/bb/7e/bbbb934247219db8299bd46dba9dd47e.jpg" width=700>

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

更节省空间的方式

``` go
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
需要注意的是，j 需要从大到小来处理。


### 0-1 背包问题升级

我们刚刚讲的背包问题，只涉及背包重量和物品重量。我们现在引入物品价值这一变量。对于一组不同重量、不同价值、不可分割的物品，我们选择将某些物品装入背包，
在满足背包最大重量限制的前提下，背包中可装入物品的总价值最大是多少呢？

#### 回溯的方式

``` java
private int maxV = Integer.MIN_VALUE; // 结果放到maxV中
private int[] items = {2，2，4，6，3};  // 物品的重量
private int[] value = {3，4，8，9，6}; // 物品的价值
private int n = 5; // 物品个数
private int w = 9; // 背包承受的最大重量
public void f(int i, int cw, int cv) { // 调用f(0, 0, 0)
  if (cw == w || i == n) { // cw==w表示装满了，i==n表示物品都考察完了
    if (cv > maxV) maxV = cv;
    return;
  }
  f(i+1, cw, cv); // 选择不装第i个物品
  if (cw + weight[i] <= w) {
    f(i+1,cw+weight[i], cv+value[i]); // 选择装第i个物品
  }
}
```

针对上面的代码，我们还是照例画出递归树。在递归树中，每个节点表示一个状态。
现在我们需要 3 个变量（i, cw, cv）来表示一个状态。其中，i 表示即将要决策第 i 个物品是否装入背包，
cw 表示当前背包中物品的总重量，cv 表示当前背包中物品的总价值。

<img src="https://static001.geekbang.org/resource/image/bf/3f/bf0aa18f367db1b8dfd392906cb5693f.jpg" width=500>

在递归树中，有几个节点的 i 和 cw 是完全相同的，比如 f(2,2,4) 和 f(2,2,3)。在背包中物品总重量一样的情况
下，f(2,2,4) 这种状态对应的物品总价值更大，我们可以舍弃 f(2,2,3) 这种状态，只需要沿着 f(2,2,4) 这条决策
路线继续往下决策就可以。

也就是说，对于 (i, cw) 相同的不同状态，那我们只需要保留 cv 值最大的那个，继续递归处理，其他状态不予考虑。

如果用回溯算法，这个问题就没法再用“备忘录”解决了。所以，我们就需要换一种思路，看看动态规划是不是更容易解决这个问题？

#### 动态规划的方式

```` go
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
````




