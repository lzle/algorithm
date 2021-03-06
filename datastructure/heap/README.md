# 堆

堆是一种特殊的树，只要满足这两点，它就是一个堆。

* 堆是一个完全二叉树；
* 堆中每一个节点的值都必须大于等于（或小于等于）其子树中每个节点的值。

对于每个节点的值都大于等于子树中每个节点值的堆，我们叫做“大顶堆”。
对于每个节点的值都小于等于子树中每个节点值的堆，我们叫做“小顶堆”。

<img src="https://static001.geekbang.org/resource/image/4c/99/4c452a1ad3b2d152daa2727d06097099.jpg" width=500>

其中第 1 个和第 2 个是大顶堆，第 3 个是小顶堆，第 4 个不是堆。同一组数据可以构建多种不同的堆。

### 如何实现一个堆?

完全二叉树适合用数组进行存储，非常节省存储空间。不需要存储左右子节点的指针，单纯地通过数组的下标，
就可以找到一个节点的左右子节点和父节点。

下标公式

> 从 1 开始，数组中下标为 i ，那左子节点的下标就是 i\*2 ，右子节点就是下标为 i\*2+1 ，父节点就是下标为 i/2。
> 从 0 开始，数组中下标为 i ，那左子节点的下标就是 2\*i+1，右子节点的下标就是 2\*i+2，父节点的下标就是 i-1/2。

### 插入元素

在堆里插入数据后，要继续满足堆的两个特性，需要对堆进行堆化操作。

堆化有两种，从下往上和从上往下，下面都是从下往上进行堆化。

下图是在数组末尾插入元素 22 时，堆化的过程。

思路如下

``` java
public class Heap {
  private int[] a; // 数组，从下标1开始存储数据
  private int n;  // 堆可以存储的最大数据个数
  private int count; // 堆中已经存储的数据个数

  public Heap(int capacity) {
    a = new int[capacity + 1];
    n = capacity;
    count = 0;
  }

  public void insert(int data) {
    if (count >= n) return; // 堆满了
    ++count;
    a[count] = data;
    int i = count;
    while (i/2 > 0 && a[i] > a[i/2]) { // 自下往上堆化
      swap(a, i, i/2); // swap()函数作用：交换下标为i和i/2的两个元素
      i = i/2;
    }
  }
 }
```

### 删除堆顶

堆顶存储的元素是最大值或最小值，当在大堆栈中，如果删除栈顶，然后用第二大元素放到堆顶，容易出现空洞。

<img src="https://static001.geekbang.org/resource/image/59/81/5916121b08da6fc0636edf1fc24b5a81.jpg" width=500>

也可以改变下思路，把最后位元素，放到堆顶，执行从上向下的堆化，可以避免空洞的出现。

<img src="https://static001.geekbang.org/resource/image/11/60/110d6f442e718f86d2a1d16095513260.jpg" width=500>

思路如下

``` java
public void removeMax() {
  if (count == 0) return -1; // 堆中没有数据
  a[1] = a[count];
  --count;
  heapify(a, count, 1);
}

private void heapify(int[] a, int n, int i) { // 自上往下堆化
  while (true) {
    int maxPos = i;
    if (i*2 <= n && a[i] < a[i*2]) maxPos = i*2;
    if (i*2+1 <= n && a[maxPos] < a[i*2+1]) maxPos = i*2+1;
    if (maxPos == i) break;
    swap(a, i, maxPos);
    i = maxPos;
  }
}
```

往堆中插入一个元素和删除堆顶元素的时间复杂度都是 O(logn)。


# 堆排序

堆排序是一种原地的、非稳定的、时间复杂度为 O(nlogn) 的排序算法。

堆排序的过程大致分解成两个大的步骤，`建堆`和`排序`。


### 建堆

第一种思路是假设一开始堆中只有 1 个元素，下标为 1 ，其他元素依次进行上面的插入操作。

第二种实现思路，跟第一种截然相反，从后往前处理数组，并且每个数据都是从上往下堆化。

<img src="https://static001.geekbang.org/resource/image/50/1e/50c1e6bc6fe68378d0a66bdccfff441e.jpg" width=500>

<img src="https://static001.geekbang.org/resource/image/aa/9d/aabb8d15b1b92d5e040895589c60419d.jpg" width=500>

实现过程

``` java
private static void buildHeap(int[] a, int n) {
  for (int i = n/2; i >= 1; --i) {
    heapify(a, n, i);
  }
}

private static void heapify(int[] a, int n, int i) {
  while (true) {
    int maxPos = i;
    if (i*2 <= n && a[i] < a[i*2]) maxPos = i*2;
    if (i*2+1 <= n && a[maxPos] < a[i*2+1]) maxPos = i*2+1;
    if (maxPos == i) break;
    swap(a, i, maxPos);
    i = maxPos;
  }
}
```

建堆过程的复杂度是多少呢？

因为叶子节点不需要堆化，所以需要堆化的节点从倒数第二层开始。每个节点堆化的过程中，需要比较和交换的节点个数，跟这个节点的高度 k 成正比。

<img src="https://static001.geekbang.org/resource/image/89/d5/899b9f1b40302c9bd5a7f77f042542d5.jpg" width=500>

高度求和公式

<img src="https://static001.geekbang.org/resource/image/f7/09/f712f8a7baade44c39edde839cefcc09.jpg" width=500>

这个公式的求解稍微有点技巧，不过我们高中应该都学过：把公式左右都乘以 2，就得到另一个公式 S2。我们将 S2 错位对齐，并且用 S2 减去 S1，可以得到 S。

<img src="https://static001.geekbang.org/resource/image/62/df/629328315decd96e349d8cb3940636df.jpg" width=500>

利用等比数列求和公式

<img src="https://static001.geekbang.org/resource/image/46/36/46ca25edc69b556b967d2c62388b7436.jpg" width=500>

因为 h=log2n，代入公式 S，就能得到 S=O(n)，所以，建堆的时间复杂度就是 O(n)。


### 排序

建堆结束之后，数组中的数据已经是按照大顶堆的特性来组织的。数组中的第一个元素就是堆顶，也就是最大的元素。
我们把它跟最后一个元素交换，那最大元素就放到了下标为 n 的位置。

这个过程有点类似上面讲的“删除堆顶元素”的操作，当堆顶元素移除之后，我们把下标为 n 的元素放到堆顶，然后
再通过堆化的方法，将剩下的 n−1 个元素重新构建成堆。堆化完成之后，我们再取堆顶的元素，放到下标是 n−1 
的位置，一直重复这个过程，直到最后堆中只剩下标为 1 的一个元素，排序工作就完成了。

<img src="https://static001.geekbang.org/resource/image/23/d1/23958f889ca48dbb8373f521708408d1.jpg" width=500>

实现过程

``` java
// n表示数据的个数，数组a中的数据从下标1到n的位置。
public static void sort(int[] a, int n) {
  buildHeap(a, n);
  int k = n;
  while (k > 1) {
    swap(a, 1, k);
    --k;
    heapify(a, k, 1);
  }
}
```

### 问题

1）为什么快速排序要比堆排序性能好？

> 第一点，堆排序数据访问的方式没有快速排序友好。数据是跳着访问的。
> 第二点，对于同样的数据，在排序过程中，堆排序算法的数据交换次数要多于快速排序。

2）建堆时，对于完全二叉树来说，下标从 2n+1 到 n 的都是叶子节点，这个结论是怎么推导出来的呢？

> 使用反证法证明，如果下标为n/2 + 1的节点不是叶子节点，即它存在子节点，按照『原理1』，它的左子节点为：2(n/2 + 1) = n + 2，大家明显可以看出，这个数字已经大于n + 1。

3）堆的应用有哪些？

> 1. 从大数量级数据中筛选出top n 条数据； 比如：从几十亿条订单日志中筛选出金额靠前的1000条数据。
> 2. 在一些场景中，会根据不同优先级来处理网络请求，此时也可以用到优先队列(用堆实现的数据结构)；比如：网络框架Volley就用了Java中PriorityBlockingQueue，当然它是线程安全的
> 3. 可以用堆来实现多路归并，从而实现有序，leetcode上也有相关的一题：Merge K Sorted Lists。

4）实现一个小顶堆、大顶堆、优先级队列？

> 建堆时间复杂度

5）实现堆排序

> [堆排序](sort.go)

6）利用优先级队列合并 K 个有序数组

> [LeetCode](https://github.com/lzle/leetcode/tree/master/note/23)
> 建立小顶堆，所有node节点放入进行堆化，每次取堆顶节点，重新建立数组。

7）求一组动态数据集合的最大 Top K

> 建大顶堆，取堆顶。









