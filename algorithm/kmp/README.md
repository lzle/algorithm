# KMP

MP 算法是根据三位作者（D.E.Knuth，J.H.Morris 和 V.R.Pratt）的名字来命名的，算法的全称是 Knuth Morris Pratt 算法，简称为 KMP 算法。
在1974年构思，发表与1977年。

KMP算法的出现对于字符串匹配意义是重大的，算法的复杂度由O(m*n)降低到O(m+n)。

算法的本质就是在于，当模式串(要查询字符串)和主串匹配时，遇到不匹配的字符时，模式串如何向后滑动的问题。

`主字串`   `a b a b a e a b a c `

`模式串` `a b a b a c`

当遇到不匹配字符 e 与 c 不匹配，此时模式串如何向后移动？

`[a b a b a e] a b a c`

`[a b a b a c]`

正常的思维是滑动一位，重新比较。如果不匹配，则再滑动一位。如果这样的话，算法的复杂度又降到O(m*n)。

`a [b a b a e a] b a c`

`* [a b a b a c]`

如何`最大化的滑动`，又`不会滑过`，是KMP的算法核心思想。

这里又涉及到两个概念：`好前缀`、`最长可匹配前缀子串`。

`好前缀`为`a b a b a`，为完全匹配的前缀字符串，如下：

<img src="https://static001.geekbang.org/resource/image/17/be/17ae3d55cf140285d1f34481e173aebe.jpg" width=500>

`最长可匹配前缀子串`为`a b a`，计算方式是在`好前缀`字符串中，前缀字符和后缀字符可匹配的最大长度，
即为`最长可匹配前缀子串`。如下。

<img src="https://static001.geekbang.org/resource/image/9e/ad/9e59c0973ffb965abdd3be5eafb492ad.jpg" width=500>

查到`最长可匹配前缀子串`，就可以计算出合理的滑动位数。滑动位数=`好前缀长度`-`最长可匹配前缀子串长度`。

`好前缀长度`为 5，`最长可匹配前缀子串长度`为 3，滑动 2 位。

`a b [a b a e a b] a c`

`* * [a b a b a c]`

重新从第一个字符串开始匹配，`好前缀`为`a b a`，`最长可匹配前缀子串`为`a`，移动位数 2。

`a b a b [a e a b a c]`

`* * * * [a b a b a c]`

接下来重新匹配，如此反复，直到匹配出结果。

KMP 算法就是利用已经匹配过的`好前缀`，从中算出合理的滑动位数。

既然计算滑动位数，需要`最长可匹配前缀子串`和`好前缀`，而`好前缀`又是在模式串当中。那么能不能预先
处理。

KMP 算法也是这么去做的，用 `next 数组`存储每个前缀（这些前缀都有可能是好前缀）的`最长可匹配前缀子串`的结尾字符下标。

<img src="https://static001.geekbang.org/resource/image/16/a8/1661d37cb190cb83d713749ff9feaea8.jpg" width=500>

以上就是 KMP 算法的思想的基本原理，理解了基本原理再来看下面的内容。

### 失效函数

`next 数组`也称为`失效函数`，如何高效的计算，是 KMP 的算法的核心内容。

照下标从小到大，依次计算 `next` 数组的值。当我们要计算 `next[i]` 的时候，前面的 `next[0]` ，`next[1]` ，`……`，`next[i-1]` 应该已经计算出来了。
利用已经计算出来的 `next` 值，我们是否可以快速推导出 `next[i]` 的值。

设 `next[i] = k` , `b` 为模式串即 `b[0,i+1]` 字符串的`最长匹配前缀字符串`为 `b[0,k+1]` 。 `i,k` 表示下标，从 `0` 开始。

例如 `a b a b c`, `b[2]` = `b[0]`, `b[0,2+1]` 字符串的`最长匹配前缀字符串`为 `b[0,0+1]` 。`next[2] = 0` 。

如果子串 b[0,k+1] 的下一个字符 b[k+2] ，与 b[0,i+1]的下一个字符 b[i+2]匹配, next[i+1] = k + 1，根据上面结论 next[i] = k，
则 `next[i+1] = next[i] + 1`。

此例中 `b[1]` 与 `b[3]` 匹配，`i = 3 , k = 1`。 则得出 `next[3] = next[2] + 1 = 1`。

根据已经推导出的 `next[i]` 的值，如果 `b[i+1]` 的值与 `b[k+1]` 的值相同，则直接可以算出 `next[i+1]` 的值。

那如果 `b[i+1]` 的值与 `b[k+1]` 的值不相同时，`next[i+1]` 的值又该如何计算哩？

接上面例子，已经推导出 `next[3] = 1`。但是下一个字符匹配时 `b[4] ≠ b[2]`。此时 `next[4] = `？，如何计算。

还是要利用 `next[3] = 1` 的结果，即 `a b a b` 的 `最长匹配前缀字符串` 为 `a b`，既然 `a b` 的下一个字符串不能与 `a b a b` 匹配，
那我们去求 `a b a b` 的 `次长匹配前缀字符串` 去试试。

问题变成了如何求出 `a b a b` 的 `次长匹配前缀字符串`，`次长匹配前缀字符串` 肯定被包含在 `最长匹配前缀字符串 a b` 内，
在前缀 `a b` 内，寻找与 `a b a b` 后缀相匹配的字符串，不就相当于在 `a b` 中寻找和 `a b` 相匹配的后缀字符串吗，也就是在 `a b` 中寻找 `最长匹配前缀字符串`。

最终问题变成，在 `b[0,4]` 中寻找 `最长匹配前缀字符串` 的 `最长匹配前缀字符串`，如果存在，查到的字符就是 `b[0,4]` 的 `好前缀`，然后取模式串的 `好前缀`
的下一个字符与 `b[4]` 进行比较。如果不存在，则查 `最长匹配前缀字符串` 的 `最长匹配前缀字符串` 的 `最长匹配前缀字符串`（晕了），直到最后无解，`next[4] = -1`。

好好消化一下。整体的思想就是假定 `next[i] = k`，如果下一个字符 `b[i+1] = b[k+1]`,则 `next[i+1] = k + 1`。
当 `b[i+1] ≠ b[k+1]` ,需要求满足 `next[i] = k` 这个公式的新的 `i、k` 值。我们的目的就是求 `b[0,i+1]` 的`好前缀`，找到`好前缀`后再匹配下一个字符 `b[i+1]`。


### 代码

``` Go
// KMP 算法，返回匹配字符串下班
// a, b分别是主串和模式串；n, m分别是主串和模式串的长度。
// 移动模式串
func kmp(a string, n int, b string, m int) int {   
    next := initNext(b, m)
    var j int
    for i := 0; i < n; i++ {
    	for j >0 && a[i] != b[j] {
    		j = next[j-1] + 1
    	}
    	if a[i] == b[j] {
    		j++
    	}
    	if j == m {
    		return i - m + 1
    	}   
    }
    return -1
}

// 初始化失效函数 next 数组
// b表示模式串，m表示模式串的长度
func initNext(b string, m int) []int {
    next := make([]int, m)
    next[0] = -1 
    // 最长匹配字符串最后下标 
    k := -1
    for i := 1; i < m; i++ {
    	for k != -1 && b[k+1] != b[i] {
    		k = next[k]
    	}
    	if b[k+1] == b[i] {
    		k++
    	}
    	next[i] = k 
    }
    return next
}

func main() {
    a := "aabbabcacdwrsg"  
    b := "sg"
    ret := kmp(a,len(a),b,len(b))
    fmt.Println(ret)
}
```





























