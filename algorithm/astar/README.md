#  A* 算法

A* 算法属于一种启发式搜索算法（Heuristically Search Algorithm），启发式算法还有其他的，例如IDA* 算法、
蚁群算法、遗传算法、模拟退火算法等。
 
回到行路线规划问题中，当时用 [Dijkstra 算法](https://github.com/lzle/algorithm/tree/master/algorithm/dijkstra)进行解决，
Dijkstra 算法有局限性，不能解决所有问题，如下。

<img src="https://static001.geekbang.org/resource/image/11/dd/11840cc13071fe2da67675338e46cadd.jpg" width=500>

从图中举的例子可以看出，尽管我们找的是从 s 到 t 的路线，但是最先被搜索到的顶点依次是 1，2，3。但是 1，2，3 三个
顶点距离t是越来越远的，搜索的方向完全不对。

之所以出现这种问题，是我们只考虑了顶点与 S 点的距离来安排队列顺序的，并没有考虑与 T 点之间的距离。

A* 算法的出去就是为了解决这个问题，在 Dijkstra 算法的基础上添加了曼哈顿距离公式。

```cgo
func hManhattan(v1 *Vertex, v2 *Vertex) int {
    return int(math.Abs(float64(v1.x - v2.x)) + math.Abs(float64(v1.y - v2.y)))
}
```

A* 算法的代码实现的主要逻辑是下面这段代码。它跟 Dijkstra 算法的代码实现，主要有 3 点区别：

* 优先级队列构建的方式不同。A* 算法是根据 f 值（也就是刚刚讲到的 f(i)=g(i)+h(i)）来构建优先级队列，而 Dijkstra 算法是根据 dist 值（也就是刚刚讲到的 g(i)）来构建优先级队列；

* A* 算法在更新顶点 dist 值的时候，会同步更新 f 值；

* 循环结束的条件也不一样。Dijkstra 算法是在终点出队列的时候才结束，A* 算法是一旦遍历到终点就结束。

```cgo
// 从顶点s到顶点t的最短路径
func astar(s int, t int, g *Graph){
    predecessor := make([]int, g.v)
    vertexes := make([]*Vertex, g.v)
    for i:=0; i<g.v; i++ {
        vertexes[i] = &Vertex{i, int(^uint(0) >> 1),int(^uint(0) >> 1),i*10,i*20}
    }
    queue := newPriorityQueue(g.v)
    inqueue := make([]bool,g.v)
    vertexes[s].dist = 0
    vertexes[s].f = 0
    queue.add(vertexes[s])
    inqueue[s] = true
    for !queue.isEmpty() {
	    minVertex := queue.poll()
        for i:=0; i<len(g.linkedList[minVertex.id]); i++ {
            edge := g.linkedList[minVertex.id][i]
            nextVertex := vertexes[edge.tid]
            if minVertex.dist + edge.w < nextVertex.dist {
                nextVertex.dist = minVertex.dist + edge.w
                nextVertex.f = nextVertex.dist + hManhattan(nextVertex,vertexes[t])
                predecessor[nextVertex.id] = minVertex.id
                if inqueue[nextVertex.id] {
                    queue.update(nextVertex)
                } else {
                    queue.add(nextVertex)
                    inqueue[nextVertex.id] = true
                }
            }
            if nextVertex.id == t {
                queue.clear()
                break   
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

A* 算法可以更加快速地找到从起点到终点的路线,但不会找的最短路线。通过曼哈顿距离公式防止路线跑偏。

