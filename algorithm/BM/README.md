## BF

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

坏字符串规则
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
































