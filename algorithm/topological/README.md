# 拓扑排序

拓扑排序算法是数据结构图中的一个经典算法

##### 可解决问题

* 检测图中环的存在

* 加载源码依赖关系


##### 如何实现加载源文件依赖关系？

可以把源文件与源文件之间的依赖关系，抽象成一个有向图。每个源文件对应图中的一个顶点，源文件之间的依赖关系就是顶点之间的边。

如果 a 先于 b 执行，也就是说 b 依赖于 a，那么就在顶点 a 和顶点 b 之间，构建一条从 a 指向 b 的边。而且，这个图不仅要是有向图，
还要是一个有向无环图，也就是不能存在像 a->b->c->a 这样的循环依赖关系。

定义图数据结构

```cgo
type Graph struct {
    v int  // 顶点的个数
    linkedList [][]int // 邻接表
}

// s先于t，边s->t
func (g *Graph)addEdge(s int, t int)  {
    g.linkedList[s] = append(g.linkedList[s],t)
}

func newGraph(v int) *Graph{
    g := new(Graph)
    g.v = v
    g.linkedList = make([][]int,v)
    return g
}
```


### Kahn算法

贪心算法思想，思路非常简单、好懂。

```cgo
func Kahn(g *Graph) {
    inDegree := make([]int, g.v)
	// 统计每个顶点入度
    for i := 0; i < g.v; i++ {
        for j := 0; j < len(g.linkedList[i]); j++ {
            w := g.linkedList[i][j]
            inDegree[w]++
        }
    }
    queue := []int{}
    for i := 0; i < g.v; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }
    for len(queue) > 0 {
        i := queue[0]
        queue = queue[:1]
        fmt.Println("->", i)
        for j := 0; j < len(g.linkedList[i]); j++ {
            w := g.linkedList[i][j]
            inDegree[w]--
            if inDegree[w] == 0 {
                queue = append(queue, i)
            }
        }
    }
}
```

> 从 Kahn 代码中可以看出来，每个顶点被访问了一次，每个边也都被访问了一次，所以，Kahn 算法的时间复杂度就是 O(V+E)
> （V 表示顶点个数，E 表示边的个数）。

### 深度优先遍历

构建逆邻表，递归处理顶点。

```cgo
func DFS(g *Graph) {
    inverse := make([][]int, g.v)
	// 构建逆邻表  边s->t表示，s依赖于t，t先于s
    for i := 0; i < g.v; i++ {
        for j := 0; j < len(g.linkedList[i]); j++ {
            w := g.linkedList[i][j]
            inverse[w] = append(inverse[w], i)
        }
    }

    var recursion func(i int, inverse [][]int, visited []bool)

    recursion = func(i int, inverse [][]int, visited []bool) {
        for j := 0; j < len(inverse[i]); j++ {
            w := inverse[i][j]
            if visited[w] {
                continue
            }
            visited[w] = true
            recursion(w, inverse, visited)
        }
        fmt.Println("->", i)
    }
    visited := make([]bool, g.v)
    for i := 0; i < g.v; i++ {
        if !visited[i] {
            recursion(i, inverse, visited)
        }
    }
}
```

>每个顶点被访问两次，每条边都被访问一次，所以时间复杂度也是 O(V+E)。