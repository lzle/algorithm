# BM

BM（Boyer-Moore）算法。它是一种非常高效的字符串匹配算法，有实验统计，它的性能是著名的KMP 算法的 3 到 4 倍。

BM 算法包含两部分，分别是 `坏字符规则` 和 `好后缀规则`。我们下面依次来看，这两个规则分别都是怎么工作的。

#### 坏字符规则

BM 算法的匹配顺序按照模式串下标从大到小的顺序，倒着匹配的。

<img src="https://static001.geekbang.org/resource/image/54/9e/540809418354024206d9989cb6cdd89e.jpg" width=500>

当我们发现某个字符没法匹配的时候。我们把这个没有匹配的字符叫作坏字符（主串中的字符）。

在模式串中并不存在这个坏字符 `c`，也就是说，只要模式串与主串的字符 `c` 重合，就无法匹配。模式串直接往后滑动三位，
重新匹配。

<img src="https://static001.geekbang.org/resource/image/a8/ca/a8d229aa217a67051fbb31b8aeb2edca.jpg" width=500>

滑动位数计算方式

把坏字符对应的模式串中的字符下标记作 `si`。如果坏字符在模式串中存在，
我们把这个坏字符在模式串中的下标记作 `xi`。如果不存在，我们把 `xi` 记作 `-1`。
那模式串往后移动的位数就等于 `si-xi`。如果坏字符在模式串里多处出现，那我们在计算 xi 的时候，
选择最靠后的那个。

<img src="https://static001.geekbang.org/resource/image/8f/2e/8f520fb9d9cec0f6ea641d4181eb432e.jpg" width=500>

不过，单纯使用坏字符规则还是不够的。因为根据 `si-xi` 计算出来的移动位数，有可能是负数，
比如主串是 `aaaaaaaaaaaaaaaa`，模式串是 `baaa`。不但不会向后滑动模式串，还有可能 `倒退`。所以，BM 算法还需要用到 `“好后缀规则”`。

针对怀字符规则做预处理，可以将模式串中的每个字符及其下标都存到散列表中。这样就可以快速找到坏字符
在模式串的位置下标了。

#### 代码
``` Go
package main

import "fmt"

const SIZE = 256

// 变量 b 是模式串，m 是模式串的长度。
// bc 表示散列表。
// 假设每个字符长度是 1 字节，用ascii码作为下标。
func initBC(b string, m int) (bc []int) {
    bc = make([]int, SIZE)
    for i := 0; i < SIZE; i++ {
    	bc[i] = -1
    }
    
    for i := 0; i < m; i++ {
    	j := int(b[i])
    	bc[j] = i
    }
    return bc
}

// bm 算法，返回匹配字符串下班
// a, b分别是主串和模式串；n, m分别是主串和模式串的长度。
// 移动主串
func bm(a string, n int, b string, m int) int{
    var (
        i int;
    	j int; 
    )
    bc := initBC(b,m)
    
    for i <= n-m {
    	for j = m-1; j >= 0; j--{
    		if a[i+j] != b[j] {
    			break
    		} 
        }
	    if j < 0 {
			return i
		}
	    i = i + (j - bc[int(a[i+j])])
    }
    return -1
}

func main() {
    a := "aabbabcacdwrsg"
    b := "sg"
    ret := bm(a,len(a),b,len(b))
    fmt.Println(ret)
}
```

#### 好后缀规则

下图，模式串和主串 `bc` 是匹配的，接下来的字符 `c` 不匹配。

<img src="https://static001.geekbang.org/resource/image/d7/8a/d78990dbcb794d1aa2cf4a3c646ae58a.jpg" width=500>

把已经匹配的 `bc` 叫作好后缀，记作{u}。在模式串的前缀中寻找，如果匹配的上，则直接滑动对齐。

<img src="https://static001.geekbang.org/resource/image/b9/63/b9785be3e91e34bbc23961f67c234b63.jpg" width=500>

不过，当模式串中不存在等于{u}的子串时，我们直接将模式串滑动到主串{u}的后面，会造成过度滑动。

<img src="https://static001.geekbang.org/resource/image/9b/70/9b3fa3d1cd9c0d0f914a9b1f518ad070.jpg" width=500>

不仅要看好后缀在模式串中，是否有另一个匹配的子串，我们还要考察`好后缀的后缀子串`，是否存在跟模式串的`前缀子串匹配`的。

<img src="https://static001.geekbang.org/resource/image/6c/f9/6caa0f61387fd2b3109fe03d803192f9.jpg" width=500>

好后缀的处理规则中最核心的内容：

* `在模式串中，查找跟好后缀匹配的另一个子串；`

* `在好后缀的后缀子串中，查找最长的、能跟模式串前缀子串匹配的后缀子串；`

预先处理每个后缀子串，对应的另一个可匹配子串的位置；接下来的构建过程是算法的核心内容。

引入最关键的变量 `suffix` 数组，`suffix` 数组的下标 `k`，表示后缀子串的长度，
下标对应的数组值存储的是，在模式串中跟好后缀{u}相匹配的子串{u*}的起始下标值，如果存在多个匹配，
存储最大的那个下标值。

<img src="https://static001.geekbang.org/resource/image/99/c2/99a6cfadf2f9a713401ba8feac2484c2.jpg" width=500>

如何记录记录模式串的 `后缀子串` 是否能匹配模式串的 `前缀子串` ？引入 `boolean` 类型的 `prefix` 数组。

<img src="https://static001.geekbang.org/resource/image/27/83/279be7d64e6254dac1a32d2f6d1a2383.jpg" width=500>


#### 代码
``` GO
package main

import "fmt"

const SIZE = 256

// 变量 b 是模式串，m 是模式串的长度。
// bc 表示散列表。
// 假设每个字符长度是 1 字节，用ascii码作为下标。
func initBC(b string, m int) (bc []int) {
	bc = make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		bc[i] = -1
	}

	for i := 0; i < m; i++ {
		j := int(b[i])
		bc[j] = i
	}
	return bc
}

func initGS(b string, m int) (suffix []int, prefix []bool) {
	var (
		j int;
		k int;
	)
	suffix = make([]int, m)
	prefix = make([]bool, m)
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}

	for i := 0; i < m-1; i++ {
		j = i
		k = 0
		for j >= 0 && b[j] == b[m-1-k] {
			k++
			suffix[k] = j
			j--
		}
		if j < 0 {
			prefix[k] = true
		}
	}
	return
}

// bm 算法，返回匹配字符串下班
// a, b分别是主串和模式串；n, m分别是主串和模式串的长度。
// 移动主串
func bm(a string, n int, b string, m int) int {
	var (
		i int;
		j int;
	)
	bc := initBC(b,m)
	suffix, prefix := initGS(b, m)

	for i <= n-m {
		for j = m - 1; j >= 0; j-- {
			if a[i+j] != b[j] {
				break
			}

		}
		if j < 0 {
			return i
		}
		x :=  j - bc[int(a[i+j])]
		y :=  moveByGS(m, j, suffix, prefix)
		if x > y {
			i = i + x
		} else {
			i = i + y
		}
	}
	return -1
}

// j 等于不匹配字符模式串下标
func moveByGS(m int, j int, suffix []int, prefix []bool) int {
	k := m - 1 - j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	} else {
		for k > 0 {
			if prefix[k] {
				return m - k
			}
			k--
		}
	}
	return m
}

func main() {
	a := "aacabbbabcabcabcdwrsg"
	b := "cabcab"
	ret := bm(a, len(a), b, len(b))
	fmt.Println(ret)
}
```





















