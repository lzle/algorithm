# 图

图分为`无向图`、`有向图`、`带权图`。

无向图中的每个点为`顶点（vertex）`，顶点直接相连接的线叫做`边（edge）`。边的条数，叫作顶点的`度（degree）`。

有向图每条边多了方向。入方向的读称为`入度（In-degree）`，出方向为`出度（Out-degree）`。

带权图每条边都有一个`权重`。

存储方式：

* 邻接矩阵存储方法

* 邻接表存储方法

涉及算法：

* 图的搜索（深度优先搜索、广度优先搜索、A*、IDA* 启发式搜索）

* 最短路径 (Floyd-Warshall算法，Dijkstra 算法，A*搜索)

* 最小生成树

* 二分图

* 拓扑排序

### 邻接矩阵

邻接矩阵是最`直观`的存储方式。

无向图中如果顶点 i 与顶点 j 之间有边，我们就将 A[i][j]和 A[j][i]标记为 1。有向图标记有方向的点为1。带权图存储的值为权重。

<img src="https://static001.geekbang.org/resource/image/62/d2/625e7493b5470e774b5aa91fb4fdb9d2.jpg" width=500>

优点：
* `邻接矩阵的存储方式简单。查询关系时非常高效`

缺点：
* `浪费空间，不适用于顶点过多的图。`

### 邻接表

顶点存储的是一条链表。

<img src="https://static001.geekbang.org/resource/image/03/94/039bc254b97bd11670cdc4bf2a8e1394.jpg" width=500>

邻接表存储起来比较节省空间，但是使用起来就比较耗时间。

链表改成`平衡二叉查找树`，实际中可以使用`红黑树`，也可以换成`跳表`、`散列表`。除此之外，有序`动态数组`。

### BFS/DFS

1）[深度优先](https://github.com/lzle/algorithm/blob/master/datastructure/graph/graph.go) :lemon:

2）[广度优先](https://github.com/lzle/algorithm/blob/master/datastructure/graph/graph.go) :lemon:


广度优先搜索和深度优先搜索是图上的两种最常用、最基本的搜索算法。比起其他高级的搜索算法，比如 A*、IDA* 等，
要简单粗暴，没有什么优化，所以，也被叫作暴力搜索算法。所以，这两种搜索算法仅适用于状态空间不大，也就是说图不大的搜索。

其中`深度优先搜索`找出来的路径，并`不是`顶点 s 到顶点 t 的`最短路径`。

在执行效率方面，深度优先和广度优先搜索的时间复杂度都是 `O(E)`，空间复杂度是 `O(V)`。


### 思考 🤔

1）如何存储微博、微信等社交网络中的好友关系？


### 解答

1） `以微博为例，除了有关注的人，还有粉丝，所以需要建立逆邻接表。将邻接表中的链表改成跳表。跳表插入、删除、查找都非常高效，
时间复杂度是 O(logn)，空间复杂度上稍高，是 O(n)。数据量太大，内存存储不下，可以使用一致性哈希进行数据分片。
另一种思路是使用外部存储，数据库对多对的关系表。`



