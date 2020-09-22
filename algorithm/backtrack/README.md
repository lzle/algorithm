# 回溯

回溯的处理思想，有点类似枚举搜索，枚举所有的解，找到满足期望的解。

当面对一个岔路口，我们先随意选一条路走，当发现这条路走不通的时候（不符合期望的解），
就回退到上一个岔路口，另选一种走法继续走。

回溯算法非常适合用递归来实现，在实现的过程中，剪枝操作是提高回溯效率的一种技巧。
利用剪枝，我们并不需要穷举搜索所有的情况，从而提高搜索效率。

很多经典的数学问题都可以用回溯算法解决，比如[数独](https://github.com/lzle/leetcode/tree/master/note/37)、八皇后、0-1 背包、图的着色、旅行商问题、[全排列](https://github.com/lzle/leetcode/tree/master/note/46)等等。

### 八皇后问题

有一个 8x8 的棋盘，希望往里放 8 个棋子（皇后），每个棋子所在的行、列、对角线都不能有另一个棋子。
你可以看我画的图，第一幅图是满足条件的一种方法，第二幅图是不满足条件的。八皇后问题就是期望找到所有
满足这种要求的放棋子方式。

<img src="https://static001.geekbang.org/resource/image/a0/f5/a0e3994319732ca77c81e0f92cc77ff5.jpg" width=500>

把问题划分为8个阶段，每个阶段对应选择一行。

核心代码如下：

```
var result [8]int

func call8Queens(row int) {
    if row == 8 {
    	printQueues()
    	os.Exit(100)
    }
    for column:=0; column<8; column++ {
    	if isOk(row,column) == true {
    	    result[row] = column
    	    call8Queens(row+1)
    	}
    }
}

func isOk(row int, column int) bool {
    left,right := column-1,column+1
    for i:=row-1; i>=0; i-- {
    	if result[i] == column {
    	    return false
    	}
    	if left >= 0 {
    	    if result[i] == left {
    		    return false
    	    }
    	}
    	if right < 8 {
    	    if result[i] == right{
    		    return false
    	    }
        }
        left--
        right++
    }
    return true
}
```

### 0-1 背包

0-1 背包问题有很多变体，我这里介绍一种比较基础的。我们有一个背包，背包总的承载重量是 Wkg。
现在我们有 n 个物品，每个物品的重量不等，并且不可分割。我们现在期望选择几件物品，装载到背包中。
在不超过背包所能装载重量的前提下，如何让背包中物品的总重量最大？

对于每个物品来说，都有两种选择，装进背包或者不装进背包。对于 n 个物品来说，总的装法就有 2^n 种，
去掉总重量超过 Wkg 的，从剩下的装法中选择总重量最接近 Wkg 的。

利用回溯，每次执行分为装与不装。

核心代码如下：

```
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
```

### 正则表达式

正则表达式中，最重要的就是通配符，通配符结合在一起，可以表达非常丰富的语义。
为了方便讲解，我假设正则表达式中只包含'\*'和'?'这两种通配符，并且对这两个通
配符的语义稍微做些改变，其中，'*'匹配任意多个（大于等于 0 个）任意字符，'?'
匹配零个或者一个任意字符。

正则表达式与字符串完全匹配

> pattern := "a*b?c"
>
>text := "accccccccbc"

核心代码如下：

```
var matched = false

func regular(pattern string, text string) bool {
    match(0, 0, pattern, text)
    return matched
}

func match(pi int, tj int, pattern string, text string){
    if matched {
	    return
    }
    if pi == len(pattern) {
    	if tj == len(text) {
    		fmt.Println()
    		matched = true
    	}
    	return
    }
    if pattern[pi] == '*' {
    	for k:=0; k<len(text)-tj; k++ {
    		match(pi+1, tj+k, pattern,text)
    	}
    } else if pattern[pi] == '?' {
    	match(pi+1, tj, pattern, text)
    	match(pi+1, tj+1, pattern,text)
    } else if tj < len(text) && pattern[pi] == text[tj] {
    	match(pi+1, tj+1, pattern,text)
    }
}
```
