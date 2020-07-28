package main

import (
	"container/list"
	"fmt"
)

// 无向图
type Graph struct {
	// 双向链表
	adj []*list.List
	v   int
}

// 初始化
func newGraph(v int) *Graph {
	g := new(Graph)
	g.v = v
	g.adj = make([]*list.List, v)
	for i := 0; i < v; i++ {
		g.adj[i] = list.New()
	}
	return g
}

// 建立关系
func (g *Graph) addEdge(s int, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

// 广度优先，从s->t
func (g *Graph) BFS(s int, t int) {
	if s == t {
		return
	}

	// 存储路径
	prev := make([]int, g.v)
	for i := range prev {
		prev[i] = -1
	}

	var queue []int
	visited := make([]bool, g.v)
	queue = append(queue, s)
	visited[s] = true
	isFound := false

	for len(queue) > 0 && !isFound {
		vertex := queue[0]
		queue = queue[1:]
		linkedList := g.adj[vertex]
		for e := linkedList.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = vertex
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}
	if isFound {
		printPrev(prev, t)
	} else {
		fmt.Printf("no path found from %d to %d", s, t)
	}
}

// 深度优先 s—>t
func (g *Graph) DFS(s int, t int) {
	prev := make([]int, g.v)
	for i := range prev {
		prev[i] = -1
	}
	visited := make([]bool, g.v)
	visited[s] = true
	g.recurse(s,t,prev,visited)
	printPrev(prev,t)
}

func (g *Graph) recurse(s int, t int, prev []int, visited []bool) bool {
	if s == t {
		return true
	}
	visited[s] = true

	linkedList := g.adj[s]
	for e := linkedList.Front(); e!= nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			if g.recurse(k,t,prev,visited){
				return true
			}
		}
	}
	return false
}


func printPrev(prev []int, t int) {
	if prev[t] == -1 {
		fmt.Printf("%d ", t)
	} else {
		printPrev(prev, prev[t])
		fmt.Printf("%d ", t)
	}
}

func main() {
	graph := newGraph(8)
	graph.addEdge(0, 1)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)
	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(3, 4)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 7)
	graph.addEdge(6, 7)

	graph.BFS(0, 7)
	fmt.Println()
	graph.BFS(1, 3)
	fmt.Println()
	graph.DFS(0, 7)
	fmt.Println()
	graph.DFS(1, 3)
	fmt.Println()
}
