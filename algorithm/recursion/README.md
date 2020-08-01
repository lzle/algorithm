## 递归

递归是一种算法，或称编程技巧。

递归把大的问题分解为小的子问题。去的过程叫“递”，回来的过程叫“归”。

满足递归的条件：

* `一个问题的解可以分解为几个子问题的解`

* `这个问题与分解之后的子问题，除了数据规模不同，求解思路完全一样`

* `存在递归终止条件`

公式：

`f(n) = f(n-1)+f(n-2)`

注意的地方：

* `递归要警惕堆栈溢出，设置递归深度。`

* `递归代码要警惕重复计算。可以用字典存储计算结果`


## 练习

1）实现斐波那契数列:

``` Go
// 可以添加map存储已经计算过的值，避免重复计算。
func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main()  {
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(10))
}
```

2）编程实现求阶乘

``` Go
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main()  {
	fmt.Println(factorial(1))
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
}
```

3) [爬梯子](https://github.com/lzle/leetcode/tree/master/note/70) :green_apple:

