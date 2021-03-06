# 贪心算法

贪心算法、分治算法、回溯算法、动态规划，它们应该是算法思想，并不是具体的算法，常用来指导我们设计具体的算法和编码等。

贪心算法有很多经典的应用：

* 霍夫曼编码

* Prim 和 Kruskal 最小生成树算法

* Dijkstra 单源最短路径算法

贪心算法解决问题的步骤,针对一组数据，我们定义了限制值和期望值，希望从中选出几个数据，在满足`限制值`的情况下，`期望值`最大。

利用局部最优解，每次选择当前情况下，在对限制值同等贡献量的情况下，选择对期望值贡献最大的数据。

实际上，用贪心算法解决问题的思路，并`不总能给出最优解`。


### 为何不是最优解

在有权图中，我们从顶点 S 开始，找一条到顶点 T 的最短路径（路径中边的权值和最小）。

<img src="https://static001.geekbang.org/resource/image/2d/42/2de91c0afb0912378c5acf32a173f642.jpg" width=500>

贪心算法的解决思路是每次选择最小权重最小的路径，求出的最短路径是 S->A->E->T。

而实际上最短路径是S->B->D->T。

贪心算法不工作的主要原因是，前面的选择，会影响后面的选择。即便我们第一步选择最优的走法（边最短），但有可能因为这一步选择，
导致后面每一步的选择都很糟糕，最终也就无缘全局最优解了。


### 分糖果

我们有 m 个糖果和 n 个孩子。糖果少，孩子多（m<n)。

每个糖果的大小不等，这 m 个糖果的大小分别是 s1，s2，s3，……，sm。
除此之外，每个孩子对糖果大小的需求也是不一样的，只有糖果的大小大于等于孩子的对糖果大小的需求的时候，孩子才得到满足。

如何分配糖果，能尽可能满足最多数量的孩子？

可以把这个问题抽象成，从 n 个孩子中，抽取一部分孩子分配糖果，让满足的孩子的个数（期望值）是最大的。这个问题的限制值就是糖果个数 m。

利用贪心算法，每次从剩下的孩子中，找出对糖果大小需求最小的，然后发给他剩下的糖果中能满足他的最小的糖果，
这样得到的分配方案，也就是满足的孩子个数最多的方案。


### 区间覆盖

假设我们有 n 个区间，区间的起始端点和结束端点分别是[l1, r1]，[l2, r2]，[l3, r3]，……，[ln, rn]。我们从这 n 个区间中选出一部分区间，
这部分区间满足两两不相交（端点相交的情况不算相交），最多能选出多少个区间呢？

<img src="https://static001.geekbang.org/resource/image/f0/cd/f0a1b7978711651d9f084d19a70805cd.jpg" width=500>

这个问题的处理思路稍微不是那么好懂，不过，我建议你最好能弄懂，因为这个处理思想在很多贪心算法问题中都有用到，比如任务调度、
教师排课等等问题。

这个问题的解决思路是这样的：我们假设这 n 个区间中最左端点是 lmin，最右端点是 rmax。
这个问题就相当于，我们选择几个不相交的区间，从左到右将[lmin, rmax]覆盖上。我们按照起始端点从小到大的顺序对这 n 个区间排序。

我们每次选择的时候，左端点跟前面的已经覆盖的区间不重合的，
右端点又尽量小的，这样可以让剩下的未覆盖区间尽可能的大，就可以放置更多的区间。这实际上就是一种贪心的选择方法。

<img src="https://static001.geekbang.org/resource/image/ef/b5/ef2d0bd8284cb6e69294566a45b0e2b5.jpg" width=500>


### 霍夫曼编码

[维基百科](https://zh.wikipedia.org/wiki/%E9%9C%8D%E5%A4%AB%E6%9B%BC%E7%BC%96%E7%A0%81)
[编码实现](https://github.com/lzle/algorithm/tree/master/algorithm/huffmancoding)

《数据结构与算法分析》书中是这样说的，"该算法是贪婪算法的原因在于，在每一阶段我们都进行一次合并而没有
进行全局的考虑。我们只是选择两颗最小的树。"

如何使用贪心算法实现霍夫曼编码？

接下来样例数据存在特殊筛选，满足`任何一个字符的编码都不是另一个的前缀`，真实情况下，不会这样，例子只是便于理解，不能作为生产使用。

假设我有一个包含 1000 个字符的文件，每个字符占 1 个 byte（1byte=8bits），存储这 1000 个字符就一共需要 8000bits，那有没有更加节省空间的存储方式呢？

通过统计分析发现，这 1000 个字符中只包含 6 种不同字符，假设它们分别是 a、b、c、d、e、f。

每个字符我们用 3 个二进制位来表示。那存储这 1000 个字符只需要 3000bits 就可以了。不过有没有更加节省空间的存储方式。

> a(000)、b(001)、c(010)、d(011)、e(100)、f(101)

霍夫曼编码就要登场了。霍夫曼编码是一种十分有效的编码方法，广泛用于数据压缩中，其压缩率通常在 20%～90% 之间。

霍夫曼编码除了考察字符个数以外，还会考虑每个字符出现的频率，根据频率的不同，选择不同长度的编码，进一步增加压缩的效率。

假设这 6 个字符出现的频率从高到低依次是 a、b、c、d、e、f。把它们编码下面这个样子，任何一个字符的编码都不是另一个的前缀，在解压缩的时候，
我们每次会读取尽可能长的可解压的二进制串，所以在解压缩的时候也不会歧义。

<img src="https://static001.geekbang.org/resource/image/83/45/83921e609c8a4dc81ca5b90c8b4cd745.jpg" width=500>

我们把每个字符看作一个节点，并且附带着把频率放到优先级队列中。我们从队列中取出频率最小的两个节点 A、B，然后新建一个节点 C，把频率设置为两个节点的频率之和，
并把这个新节点 C 作为节点 A、B 的父节点。最后再把 C 节点放入到优先级队列中。重复这个过程，直到队列中没有数据。

<img src="https://static001.geekbang.org/resource/image/7b/7a/7b6a08e7df45eac66820b959c64f877a.jpg" width=500>

现在，我们给每一条边加上画一个权值，指向左子节点的边我们统统标记为 0，指向右子节点的边，我们统统标记为 1，那从根节点到叶节点的路径就是叶节点对应字符的霍夫曼编码。

<img src="https://static001.geekbang.org/resource/image/cc/ed/ccf15d048be005924a409574dce143ed.jpg" width=500>



