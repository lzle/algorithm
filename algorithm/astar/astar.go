package astar

import (
	"fmt"
	"math"
)

type Graph struct {
	v          int       // 顶点的个数
	linkedList [][]*Edge // 邻接表
}

type Edge struct {
	sid int
	tid int
	w   int // 权重
}

type Vertex struct {
	id   int // 顶点编码ID
	dist int // 从起始顶点到这个顶点的距离
	f	int  // 新增：f = dist + 曼哈顿距离
	x 	int  // 横坐标
	y 	int  // 纵坐标
}

// 添加边
func (g *Graph) addEdge(s int, t int, w int) {
	g.linkedList[s] = append(g.linkedList[s],
		&Edge{s, t, w})
}

func newGraph(v int) *Graph {
	g := new(Graph)
	g.v = v
	g.linkedList = make([][]*Edge, v)
	return g
}

// 小顶堆
type PriorityQueue struct {
	nodes []*Vertex
	count int
}

// 获取堆顶点
func (q *PriorityQueue) poll() *Vertex {
	return nil
}

// 添加
func (q *PriorityQueue) add(v *Vertex) {

}

// 添加
func (q *PriorityQueue) clear() {

}

// 更新结点的值，并且从下往上堆化，重新符合堆的定义。
// 时间复杂度O(logn)。
func (q *PriorityQueue) update(v *Vertex) {

}

func (q *PriorityQueue) isEmpty() bool {
	if len(q.nodes) == 0 {
		return true
	}
	return false
}

func newPriorityQueue(v int) *PriorityQueue {
	q := new(PriorityQueue)
	q.count = v
	q.nodes = make([]*Vertex, v+1)
	return q
}

func hManhattan(v1 *Vertex, v2 *Vertex) int {
	return int(math.Abs(float64(v1.x - v2.x)) + math.Abs(float64(v1.y - v2.y)))
}


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











