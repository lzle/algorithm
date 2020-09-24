package topological

import "fmt"

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
