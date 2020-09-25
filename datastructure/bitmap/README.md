## 位图

我们有 1 千万个整数，整数的范围在 1 到 1 亿之间。如何快速查找某个整数是否在这 1 千万个整数中呢？

除了可以通过散列表进行存储以为，还可以使用一种特殊的散列表：位图。

申请 1 亿大小、布尔数据类型的数组，将 1 千万个整数作为数组下标，对应的值设置为true。
当查询是，只需要查看对应的数组下标是否为ture即可，查询复杂度为 O(1)。布尔类型的占用空间
为 1 字节，表示 true 和 false 两个值，没必要用布尔类型，只需要用二进制位(bit)即可。
具体实现如下。

```cgo
type BitMap struct {
    nbits int
    bytes []int
}

func (bm *BitMap) Set(k int) {
    if k > bm.nbits {
        return
    }
    byteIndex := k / 8
    bitIndex := k % 8
    bm.bytes[byteIndex] =  bm.bytes[byteIndex] | (1 << bitIndex)
}

func (bm *BitMap) Get(k int) bool {
    if k > bm.nbits {
        return false
    }
    byteIndex := k / 8
    bitIndex := k % 8
    return (bm.bytes[byteIndex] & (1 << bitIndex)) != 0
}

func newBitMap(n int) *BitMap {
    bm := new(BitMap)
    bm.nbits = n
    bm.bytes = make([]int, n/8+1)
    return bm
}
```

回到刚才的例子，如果使用散列表进行存储，1 千万个数据，数据是 32 位的整型数，也就是需要 4 个字节的存储空间，那总共至少
需要 40MB 的存储空间，实际上占用空间更大(散列冲突)。而如果通过位图的话，数字范围在 1 到 1 亿之间，只需要 1 亿个二进制位，
也就是 12MB 左右的存储空间就够了。


## 布隆过滤器

上面例子中，假如数字范围在 1 到 10 亿之间呢，那就需要 10 亿个二进制位，120MB 的存储空间，比散列表更大，位图的优势完全不存在。

为了解决这个问题，布隆过滤器出场。

解决问题的思想是，布隆过滤器使用哈希函数，使得 10 亿数值的数据还落在 1 亿个二进制位图区间内，位图占用空间依然是 12MB。

不过哈希函数肯定会造成冲突，一亿零一和 1 两个数字进行取模求余后的值一样，这样就无法进行区分是一亿零一 还是 1 了。

布隆过滤器通过 K 个哈希函数，得到 K 个 Y 值，把不同 Y 值对应的位图全部设置为 true 。查询时，如果值全部为 true 则表明存在，
其中有一个不为true，则不存在。

<img src="https://static001.geekbang.org/resource/image/94/ae/94630c1c3b7657f560a1825bd9d02cae.jpg" width=500>

经过 K 个哈希函数处理之后，K 个哈希值都相同的概率就非常低了，但是，这种处理方式又带来了新的问题，那就是容易误判。

<img src="https://static001.geekbang.org/resource/image/d0/1a/d0a3326ef0037f64102163209301aa1a.jpg" width=500>


布隆过滤器的误判有一个特点，那就是，它只会对`存在的情况`有误判。如果某个数字经过布隆过滤器判断不存在，
那说明这个数字`真的不存在`，不会发生误判。对于某些场景，这是可以进行接受的。