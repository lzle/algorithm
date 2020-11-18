# 快速排序

快排的思想是这样的：如果要排序数组中下标从 p 到 r 之间的一组数据，
我们选择 p 到 r 之间的任意一个数据作为 pivot（分区点）。
 
把小于 pivot 的数据放左边，大于 pivot 的数据放右边。经过分区，数据分为了三个部分。
 
<img src="https://static001.geekbang.org/resource/image/4d/81/4d892c3a2e08a17f16097d07ea088a81.jpg" width=500>

主体逻辑如下

```
// 快速排序，A是数组，n表示数组的大小
quick_sort(A, n) {
  quick_sort_c(A, 0, n-1)
}
// 快速排序递归函数，p,r为下标
quick_sort_c(A, p, r) {
  if p >= r then return
  
  q = partition(A, p, r) // 获取分区点
  quick_sort_c(A, p, q-1)
  quick_sort_c(A, q+1, r)
}
```

快排的核心在于 partition 分区函数。partition 实现原地排序，空间复杂度为 O(1)。

原地分区函数的实现思路非常巧妙

```
partition(A, p, r) {
  pivot := A[r]
  i := p
  for j := p to r-1 do {
    if A[j] < pivot {
      swap A[i] with A[j]
      i := i+1
    }
  }
  swap A[i] with A[r]
  return i

```

图示来展示分区的整个过程

<img src="https://static001.geekbang.org/resource/image/08/e7/086002d67995e4769473b3f50dd96de7.jpg" width=500>

### 总结

快排是一种原地、不稳定的排序算法。时间复杂度为O(nlogn)。

快速排序算法虽然最坏情况下的时间复杂度是 O(n2)，但是平均情况下时间复杂度都是 O(nlogn)。不仅如此，快速排序算法时间复杂度退化到 O(n2) 的概率非常小，
我们可以通过合理地选择 pivot 来避免这种情况。

### 问题

1）如何在 O(n) 的时间复杂度内查找一个无序数组中的第 K 大元素？ 

> 思路 partition 函数返回下标。

2）为何是不稳定的？

> 如果数组中有两个相同的元素，比如序列 6，8，7，6，3，5，9，4，在经过第一次分区操作之后，两个 6 的相对先后顺序就会改变。所以，快速排序并不是一个稳定的排序算法。
