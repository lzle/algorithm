package topological

import "fmt"

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
