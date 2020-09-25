#  Dijkstra 算法

Dijkstra 算法是最短路径算法中最经典的一个。
 
用于解决地图软件中从起始地址到结束地址中，最优出现路线的选择。

#### 解决思想

针对地图道路进行建模，把岔路口看作顶点，道路看作顶点之间的边，路长看作边的权重，单、双行道看作有向边，这样整个地图被抽象成有向有权图。

#### 核心代码。

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

#### 步骤

<img src="https://static001.geekbang.org/resource/image/e2/a9/e20907173c458fac741e556c947bb9a9.jpg" width=500>

Dijkstra 算法的时间复杂度 O(E*logV) 。利用的是动态规划的思想。

最短路径算法还有很多，比如 `Bellford 算法`、`Floyd 算法`等等。

真实的地图软件的路径规划，要比这个复杂很多。而且，比起 Dijkstra 算法，地图软件用的更多的是类似` A* 的启发式搜索算法`。
