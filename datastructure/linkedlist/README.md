# 链表

相较于数组，链表是`不需要`一块连续的内存空间，它通过`“指针”`将一组零散的内存块`串联`起来使用

常见的链表结构：`单链表`、`双链表`、`循环链表`

单链表：`结点存储后继指针和数据，尾结点指向Null`

双向链表：`结点存储前驱指针、后继指针和数据，尾结点指向Null`

循环链表：`尾结点后继指针指向头结点`

## 优缺点

* `与数组相比，需要额外的空间存储指针信息。`

* `内存申请灵活，不需要连续的地址空间。但容易操作内存碎片。`

* `数组可以利用 CPU 缓存机制，预读数组中的数据。`

* `更适合插入、删除频繁的场景。`

## 时间复杂度

查询  `O(n)`

插入  `理论上 O(1)`

删除  `理论上 O(1)`

为什么说插入、删除`理论`上时间复杂度是` O(1)`呢？

`在实际的软件开发中，从链表中删除一个数据无外乎这两种情况：`

* `删除结点中“值等于某个给定值”的结点；`

* `删除给定指针指向的结点。`

`对于第一种情况，单链表还是双向链表，时间复杂度都是 O(n)。需要查找操作。`

`对于第二种情况，单链表需要寻找到前驱结点，所有事件复杂度为 O(n)。双向链表结点已经存放有前驱结点的指针，时间复杂度为 O(1)。`

思考 🤔

`对于插入，插入也分为在当前节点前、节点后，以及添加到index位置下。`

`这三种情况下，单链表、双向链表时间复杂度又各是多少？`

## 练习

1）[实现单链表](https://github.com/lzle/algorithm/blob/master/datastructure/linkedlist/single.go) :apple:

2）[实现双链表]() :apple:

3）[实现循环链表]() :apple:

4）[单链表反转](https://github.com/lzle/leetcode/tree/master/note/206) green_apple:

5）[链表中环的检测](https://github.com/lzle/leetcode/tree/master/note/141) :green_apple:

6）[两个有序的链表合并](https://github.com/lzle/leetcode/tree/master/note/21)  :green_apple:

7）[删除链表倒数第 n 个结点](https://github.com/lzle/leetcode/tree/master/note/19)  :lemon:

8）[求链表的中间结点](https://github.com/lzle/leetcode/tree/master/note/876)  :green_apple:

9）[合并 k 个排序链表](https://github.com/lzle/leetcode/tree/master/note/23)  :apple:


## 补充

1）LRU `最近最少使用策略（Least Recently Used）`

2）[约瑟夫问题](https://zh.wikipedia.org/wiki/%E7%BA%A6%E7%91%9F%E5%A4%AB%E6%96%AF%E9%97%AE%E9%A2%98)

3） `CPU在从内存读取数据的时候，会先把读取到的数据加载到CPU的缓存中。而CPU每次从内存读取数据并不是只读取那个特定要访问的地址，而是读取一个数据块(这个大小我不太确定。。)并保存到CPU缓存中，然后下次访问内存数据的时候就会先从CPU缓存开始查找，如果找到就不需要再从内存中取。这样就实现了比内存访问速度更快的机制，也就是CPU缓存存在的意义:为了弥补内存访问速度过慢与CPU执行速度快之间的差异而引入。`
