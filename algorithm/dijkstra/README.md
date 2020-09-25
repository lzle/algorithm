#  Dijkstra 算法

像 Google 地图、百度地图、高德地图这样的地图软件，输入起始、结束地址，地图就会给你规划一条最优出行路线，这是如何实现的？

想要解决这个问题，有一个非常经典的算法，最短路径算法，更加准确地说，是单源最短路径算法（一个顶点到一个顶点）。提到最短路径算法，
最出名的莫过于 Dijkstra 算法了。

首先要进行建模，我们把每个岔路口看作一个顶点，岔路口与岔路口之间的路看作一条边，路的长度就是边的权重。如果路是单行道，我们就
在两个顶点之间画一条有向边；如果路是双行道，我们就在两个顶点之间画两条方向不同的边。整个地图就被抽象成一个有向有权图。

原理稍微有点儿复杂，看核心代码。

```cgo
// 从顶点s到顶点t的最短路径
func dijkstra(s int, t int, g *Graph){
    predecessor := make([]int, g.v)
    vertexes := make([]*Vertex, g.v)
    for i:=0; i<g.v; i++ {
        vertexes[i] = &Vertex{i, int(^uint(0) >> 1)}
    }
    queue := newPriorityQueue(g.v)
    inqueue := make([]bool,g.v)
    vertexes[s].dist = 0
    queue.add(vertexes[s])
    inqueue[s] = true
    for !queue.isEmpty() {
        minVertex := queue.poll()
        if minVertex.id == t {
            break
        }
        for i:=0; i<len(g.linkedList[minVertex.id]); i++ {
            edge := g.linkedList[minVertex.id][i]
            nextVertex := vertexes[edge.tid]
            if minVertex.dist + edge.w < nextVertex.dist {
                nextVertex.dist = minVertex.dist + edge.w
                predecessor[nextVertex.id] = minVertex.id
                if inqueue[nextVertex.id] {
                    queue.update(nextVertex)
                } else {
                    queue.add(nextVertex)
                    inqueue[nextVertex.id] = true
                }
            }
        }
    }
    var f func(s int, t int)
    f = func(s int, t int) {
        if s != t {
            f(s, predecessor[t])
        }
        fmt.Println(t)
    }
    f(s,t)
}
```

根据步骤绘图。

<img src="https://static001.geekbang.org/resource/image/e2/a9/e20907173c458fac741e556c947bb9a9.jpg" width=500>

Dijkstra 算法的时间复杂度 O(E*logV) 。利用的是动态规划的思想。

最短路径算法还有很多，比如 Bellford 算法、Floyd 算法等等。

真实的地图软件的路径规划，要比这个复杂很多。而且，比起 Dijkstra 算法，地图软件用的更多的是类似 A* 的启发式搜索算法。